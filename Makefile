# Load local environment variables if .env exists
-include .env
export $(shell sed 's/=.*//' .env 2>/dev/null || true)

up:
	docker-compose up

up-d:
	docker-compose up -d

down:
	docker-compose down

build:
	docker-compose build

build-nc:
	docker-compose build --no-cache

migrate-up:
	@echo "Current directory: $(PWD)"
	docker run --rm $(if $(DOCKER_NETWORK),--network $(DOCKER_NETWORK)) \
		-v $(PWD)/migrations:/app migrate/migrate \
		-path=/app -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=require" up
