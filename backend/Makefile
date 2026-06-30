.PHONY: help dev build test lint clean sqlc swagger migrate-up migrate-down docker-up docker-down

# ============================================
# Hermes TodoList — Backend Makefile
# ============================================

help: ## Show this help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | \
		awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-20s\033[0m %s\n", $$1, $$2}'

# ---- Development ----
dev: ## Start with hot reload (requires air)
	air

run: ## Run without hot reload
	go run ./cmd/server

# ---- Build ----
build: ## Build production binary
	CGO_ENABLED=0 go build -ldflags="-s -w" -o bin/server ./cmd/server

# ---- Test ----
test: ## Run all tests
	go test ./... -v -race -count=1

test-cover: ## Run tests with coverage
	go test ./... -v -race -coverprofile=coverage.out -covermode=atomic
	go tool cover -html=coverage.out -o coverage.html

test-short: ## Run tests (skip integration)
	go test ./... -short -v

# ---- Lint ----
lint: ## Run golangci-lint
	golangci-lint run --timeout=5m ./...

# ---- Code Generation ----
sqlc: ## Generate sqlc Go code
	sqlc generate

swagger: ## Generate Swagger docs
	swag init -g cmd/server/main.go -o docs/swagger

# ---- Database ----
migrate-up: ## Run database migrations
	go run ./cmd/migrate up

migrate-down: ## Rollback last migration
	go run ./cmd/migrate down

migrate-create: ## Create new migration (usage: make migrate-create NAME=add_something)
	go run ./cmd/migrate create $(NAME)

# ---- Docker ----
docker-build: ## Build Docker image
	docker build -t hermes-todolist-api .

docker-up: ## Start with Docker Compose
	docker compose up -d

docker-down: ## Stop Docker Compose
	docker compose down

docker-logs: ## Tail Docker logs
	docker compose logs -f api

# ---- Clean ----
clean: ## Clean build artifacts
	rm -rf bin/ tmp/ coverage.out coverage.html
