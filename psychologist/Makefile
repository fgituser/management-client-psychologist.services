MODULE = $(shell go list -m)
CONFIG_FILE ?= ./configs/local.yaml
APP_DSN ?= $(shell sed -n 's/^dsn:[[:space:]]*"\(.*\)"/\1/p' $(CONFIG_FILE))
MIGRATE := docker run -v $(shell pwd)/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database "$(APP_DSN)"

.PHONY: build
build:  ## build the API server binary
	@echo $(MODULE)/
	GO111MODULE=on CGO_ENABLED=1 go build -o service $(MODULE)/cmd/service	

.PHONY: db-start
db-start: ## start the database server
	@mkdir -p testdata/postgres
	docker run --rm --name postgres_psychologist -v $(shell pwd)/testdata:/testdata \
		-v $(shell pwd)/testdata/postgres:/var/lib/postgresql/data \
		-e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=psychologist -d -p 5439:5432 postgres

.PHONY: db-stop
db-stop: ## stop the database server
	docker stop postgres_psychologist

.PHONY: migrate
migrate: ## run all new database migrations
	@echo "Running all new database migrations..."
	@echo "$(APP_DSN)"
	@$(MIGRATE) up

.PHONY: migrate-create-init
migrate-create-init:
	@$(MIGRATE) create -ext sql -dir /migrations init

.PHONY: migrate-create-clients
migrate-create-clients:
	@$(MIGRATE) create -ext sql -dir /migrations clients

.PHONY: migrate-create-schedule
migrate-create-schedule:
	@$(MIGRATE) create -ext sql -dir /migrations schedule

.PHONY: migrate-create-calendar
migrate-create-calendar:
	@$(MIGRATE) create -ext sql -dir /migrations calendar

.PHONY: migrate-create-hours
migrate-create-hours:
	@$(MIGRATE) create -ext sql -dir /migrations hours

.PHONY: migrate-create-employment
migrate-create-employment:
	@$(MIGRATE) create -ext sql -dir /migrations employment

.PHONY: migrate-create-employee
migrate-create-employee:
	@$(MIGRATE) create -ext sql -dir /migrations employee

.PHONY: migrate-create-cancellation_employment
migrate-create-cancellation_employment:
	@$(MIGRATE) create -ext sql -dir /migrations cancellation_employment




.PHONY: testdata_add_case
testdata_add_case: ## populate the database with test data
	#make migrate-reset
	@echo "Populating test data..."
	@docker exec -it postgres psql "$(APP_DSN)" -f /testdata/testdata_add_case.sql

.PHONY: migrate-reset
migrate-reset: ## reset database and re-run all migrations
	@echo "Resetting database..."
	@$(MIGRATE) drop
	@echo "Running all database migrations..."
	@$(MIGRATE) up

.PHONY: migrate-down
migrate-down: ## revert database to the last migration step
	@echo "Reverting database to the last migration step..."
	@$(MIGRATE) down 1