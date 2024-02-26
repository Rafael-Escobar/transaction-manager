
.PHONY: run-local
run-local: ## Run the application locally
	@go run main.go server

PHONY: migrate-up
migrate-up:
	./scripts/migrate.sh $(ENV) up

.PHONY: migrate-create
migrate-create:
	./scripts/migrate-create.sh $(name)

.PHONY: migrate-down
migrate-down:
	./scripts/migrate.sh $(ENV) down

.PHONY: compose
compose:
	docker-compose -p registry_service -f "infra/docker/docker-compose.yml" $(c)

.PHONY: dev/up
dev/up:
	make compose c="up"

.PHONY: dev/up/d
dev/up/d:
	make compose c="up -d"

.PHONY: dev/restart
dev/restart:
	make compose c="restart app"

.PHONY: dev/down
dev/down:
	make compose c="down"
