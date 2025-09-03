FROM golang:1.25-alpine AS build-stage

WORKDIR /

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /govies ./cmd/web

FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /govies /govies

EXPOSE 4000

ENTRYPOINT ["/govies"]