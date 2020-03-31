CONFIG_FILE ?= ./configs/local.yml
APP_DSN ?= $(shell sed -n 's/^dsn:[[:space:]]*"\(.*\)"/\1/p' $(CONFIG_FILE))
MIGRATE := docker run -v $(shell pwd)/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database "$(APP_DSN)"
	

.PHONY: db-start
db-start: ## start the database server
	@mkdir -p testdata/postgres
	docker run --rm --name postgres -v $(shell pwd)/testdata:/testdata \
		-v $(shell pwd)/testdata/postgres:/var/lib/postgresql/data \
		-e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=psychologist -d -p 5432:5432 postgres

.PHONY: db-stop
db-stop: ## stop the database server
	docker stop postgres

.PHONY: migrate
migrate: ## run all new database migrations
	@echo "Running all new database migrations..."
	@echo "$(APP_DSN)"
	@$(MIGRATE) up

.PHONY: migrate-create-clients
migrate-create-clients:
	@$(MIGRATE) create -ext sql -dir /migrations clients

.PHONY: migrate-create-schedule
migrate-create-schedule:
	@$(MIGRATE) create -ext sql -dir /migrations schedule

.PHONY: migrate-create-calendar
migrate-create-calendar:
	@$(MIGRATE) create -ext sql -dir /migrations calendar



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