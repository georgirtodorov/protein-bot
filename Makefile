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
	docker run --rm -it \
		-v $(PWD):/app \
		-w /app \
		--network protein-bot_default \
		--name protein-bot-dev \
		-e DB_USER=root \
		-e DB_PASSWORD=toor \
		-e DB_HOST=db \
		-e DB_PORT=5432 \
		-e DB_NAME=protein_bot \
		-e PORT=8080 \
		-p 8080:8080 \
		golang:1.25-alpine go run ./cmd/main.go

migrate-up:
	@echo "Current directory: $(PWD)"
	docker run --rm $(if $(DOCKER_NETWORK),--network $(DOCKER_NETWORK)) \
		-v $(PWD)/migrations:/app migrate/migrate \
		-path=/app -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=require" up
