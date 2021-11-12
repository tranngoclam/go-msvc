db-up:
	@docker compose -f build/docker-compose.yaml -up --build -d

db-down:
	@docker compose -f build/docker-compose.yaml down -v --rmi local

test:
	@go test ./... -race -v

.PHONY: db-up db-down
