COMPOSE_PROFILE := go-msvc
COMPOSE_FILE := build/docker-compose.yaml
IMAGE_TAG := go-msvc

up:
	@docker compose -f $(COMPOSE_FILE) -p $(COMPOSE_PROFILE) up --build -d

down:
	@docker compose -f $(COMPOSE_FILE) -p $(COMPOSE_PROFILE) down -v --rmi local

ps:
	@docker compose -f $(COMPOSE_FILE) -p $(COMPOSE_PROFILE) ps

fmt:
	@go fmt ./...

test:
	@go clean -testcache
	@go test `go list ./...` -cover -p 1

build:
	@docker build -f build/Dockerfile -t $(IMAGE_TAG) .

.PHONY: up down ps fmt test build
