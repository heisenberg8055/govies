package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/heisenberg8055/govies/internal/realip"
	"golang.org/x/time/rate"
)

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				w.Header().Set("Connection", "Close")
				app.serverErrorResponse(w, r, fmt.Errorf("%s", err))
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func (app *application) rateLimit(next http.Handler) http.Handler {
	switch {
	case app.config.limiter.enabled:

		type client struct {
			limiter  *rate.Limiter
			lastSeen time.Time
		}
		var counter = struct {
			sync.Mutex
			m map[string]*client
		}{m: map[string]*client{}}

		// cleaning up inactive ip's every 3 mins
		go func() {
			for {
				time.Sleep(time.Minute)
				counter.Lock()
				for ip, client := range counter.m {
					if time.Since(client.lastSeen) > 3*time.Minute {
						delete(counter.m, ip)
					}
				}
				counter.Unlock()
			}
		}()

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ip := realip.FromRequest(r)
			counter.Lock()
			if _, found := counter.m[ip]; !found {
				counter.m[ip] = &client{limiter: rate.NewLimiter(2, 4)}
			}
			counter.m[ip].lastSeen = time.Now()
			if !counter.m[ip].limiter.Allow() {
				counter.Unlock()
				app.rateLimitExceededResponse(w, r)
				return
			}
			counter.Unlock()
			next.ServeHTTP(w, r)
		})
	default:
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			next.ServeHTTP(w, r)
		})
	}
}
