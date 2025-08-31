package main

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/heisenberg8055/govies/internal/data"
	"github.com/heisenberg8055/govies/internal/realip"
	"github.com/heisenberg8055/govies/internal/validator"
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

func (app *application) authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Vary", "Authorization")

		authorizationHeader := r.Header.Get("Authorization")

		if authorizationHeader == "" {
			r = app.contextSetUser(r, data.AnonymousUser)
			next.ServeHTTP(w, r)
			return
		}

		headerParts := strings.Split(authorizationHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			app.invalidAuthenticationTokenResponse(w, r)
			return
		}
		token := headerParts[1]

		v := validator.New()

		if data.ValidateTokenPlainText(v, token); !v.Valid() {
			app.invalidAuthenticationTokenResponse(w, r)
			return
		}

		user, err := app.models.Users.GetForToken(data.ScopeAuthentication, token)
		if err != nil {
			switch {
			case errors.Is(err, data.ErrRecordNotFound):
				app.invalidAuthenticationTokenResponse(w, r)
			default:
				app.serverErrorResponse(w, r, err)
			}
			return
		}
		r = app.contextSetUser(r, user)
		next.ServeHTTP(w, r)
	})
}
