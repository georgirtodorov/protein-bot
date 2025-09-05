# Load environment variables from .env
include .env
export $(shell sed 's/=.*//' .env)

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
	docker run --rm -v $(PWD)/migrations:/app --network protein-bot_default migrate/migrate \
		-path=/app -database "postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" up
