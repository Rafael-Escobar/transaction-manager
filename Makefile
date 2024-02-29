# SHELL=/bin/bash
# .ONESHELL:
# .DEFAULT_GOAL := all

MOCK_DST:=./tests/mocks
MOCK_INTERFACES_SOURCES+=TransactionRepository
MOCK_INTERFACES_SOURCES+=AccountRepository
MOCK_INTERFACES_SOURCES+=OperationTypeRepository

EMPTY :=
SPACE := $(EMPTY) $(EMPTY)
MOCKERY_NAME_ARG:="$(subst $(SPACE),|,$(MOCK_INTERFACES_SOURCES))"
MOCK_TARGETS:=$(patsubst %,$(MOCK_DST)/%.go,$(MOCK_INTERFACES_SOURCES))

.PHONY: run-local
run-local: ## Run the application locally
	@go run main.go server

.PHONY: mock-all
mock-all:
	echo $(MOCKERY_NAME_ARG) && mockery --recursive --output="$(MOCK_DST)" --name=$(MOCKERY_NAME_ARG) --dir ./internal

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
