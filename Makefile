COMPOSE_PROFILE := go-msvc
COMPOSE_FILE := build/docker-compose.yaml

up:
	@docker compose -f $(COMPOSE_FILE) -p $(COMPOSE_PROFILE) up --build -d

down:
	@docker compose -f $(COMPOSE_FILE) -p $(COMPOSE_PROFILE) down -v --rmi local

ps:
	@docker compose -f $(COMPOSE_FILE) -p $(COMPOSE_PROFILE) ps

test:
	@go clean -testcache
	@go test `go list ./...` -cover -p 1

.PHONY: up down ps test
