# Load local environment variables if .env exists
-include .env
export $(shell sed 's/=.*//' .env 2>/dev/null || true)

up:
	docker-compose up

up-d:
	docker-compose up -d

down:
	@echo "Stopping dev environment"
	docker stop protein-bot-dev || true
	docker-compose down

build:
	docker-compose build

build-nc:
	docker-compose build --no-cache

dev:
	@echo "Running locally with Go in Docker"
	docker-compose up -d db pgadmin
	docker run --rm -it\
		-v $(PWD):/app \
		-w /app/backend \
		--network $(DOCKER_NETWORK) \
		--name protein-bot-dev \
		-e DB_USER=$(DB_USER) \
		-e DB_PASSWORD=$(DB_PASSWORD) \
		-e DB_HOST=$(DB_HOST) \
		-e DB_PORT=$(DB_PORT) \
		-e DB_NAME=$(DB_NAME) \
		-e PORT=$(APP_PORT) \
		-e DB_SSLMODE=$(DB_SSLMODE) \
		-p $(APP_PORT):$(APP_PORT) \
		golang:1.25-alpine go run ./cmd/main.go

dev-d:
	@echo "Running locally with Go in Docker"
	docker-compose up -d db pgadmin
	docker run --rm -d\
		-v $(PWD):/app \
		-w /app/backend \
		--network $(DOCKER_NETWORK) \
		--name protein-bot-dev \
		-e DB_USER=$(DB_USER) \
		-e DB_PASSWORD=$(DB_PASSWORD) \
		-e DB_HOST=$(DB_HOST) \
		-e DB_PORT=$(DB_PORT) \
		-e DB_NAME=$(DB_NAME) \
		-e PORT=$(APP_PORT) \
		-e DB_SSLMODE=$(DB_SSLMODE) \
		-p $(APP_PORT):$(APP_PORT) \
		golang:1.25-alpine go run ./cmd/main.go

migrate-up:
	@echo "Current directory: $(PWD)"
	@echo "Using DB_SSLMODE=$(DB_SSLMODE)"
	docker run --rm $(if $(DOCKER_NETWORK),--network $(DOCKER_NETWORK)) \
		-v $(PWD)/backend/migrations:/app migrate/migrate \
		-path=/app -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=${DB_SSLMODE}" up

test:
	docker run --rm \
		-v $(PWD)/backend:/app/backend \
		-w /app/backend \
		golang:1.22-alpine \
		go test -v ./...