.PHONY: help build-docs serve-docs test-syntax build-server start stop test

help:
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

build-docs: ## builds swagger documentation file
	swagger generate spec -o ./docs/swagger.json --scan-models

serve-docs: build-docs ## builds then serves swagger documentation file
	swagger serve ./docs/swagger.json

test-syntax: ## runs syntax testing
	golangci-lint run ./...

build-server: ## runs backend services 
	docker-compose -f docker-compose.yml run --rm start_dependencies

run:  ## runs this service
	go run cmd/main.go

start: build-server run ## runs backend services and this service

stop:  ## stops all services
	docker-compose down

test: ## runs testing
	go test -v --coverprofile=coverage.out ./...