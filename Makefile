export GO111MODULE=on
# update app name. this is the name of binary
APP=govies
APP_EXECUTABLE="./$(APP)"
# ALL_PACKAGES=$(shell go list ./... | grep -v /vendor)
# SHELL := /bin/bash # Use bash syntax

# Optional if you need DB and migration commands
# DB_HOST=$(shell cat config/application.yml | grep -m 1 -i HOST | cut -d ":" -f2)
# DB_NAME=$(shell cat config/application.yml | grep -w -i NAME  | cut -d ":" -f2)
# DB_USER=$(shell cat config/application.yml | grep -i USERNAME | cut -d ":" -f2)

# Optional colors to beautify output
GREEN  := $(shell tput -Txterm setaf 2)
YELLOW := $(shell tput -Txterm setaf 3)
WHITE  := $(shell tput -Txterm setaf 7)
CYAN   := $(shell tput -Txterm setaf 6)
RESET  := $(shell tput -Txterm sgr0)

## Quality
audit: ## runs code quality checks
	make fmt
	make vet
	make static-check
	make govulncheck
vet: ## go vet
	go vet ./...

fmt: ## runs go formatter
	go fmt ./...

static-check:
	go tool staticcheck ./...

govulncheck:
	go tool govulncheck

tidy: ## runs tidy to fix go.mod dependencies
	go mod tidy
	go mod verify
	go mod vendor

## Test
test: ## runs tests and create generates coverage report
	make tidy
	go test -v -timeout 10m ./... -coverprofile=coverage.out -json > report.json

coverage: ## displays test coverage report in html mode
	make test
	go tool cover -html=coverage.out

## Build
build: ## build the go application
	go build -o ${APP} ./cmd/govies

run: ## runs the go binary. use additional options if required.
	make build
	chmod +x $(APP_EXECUTABLE)
	$(APP_EXECUTABLE)

## Docker compose
docker-build:
	docker compose build --no-cache

docker-run:
	make docker-build
	docker compose up