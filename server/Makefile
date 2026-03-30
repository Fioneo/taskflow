include .env
export 
export PROJECT_ROOT=${shell pwd}
env-up:
	@docker compose up -d db
env-down:
	@docker compose down db
env-cleanup:
	@read -p "Clean up all environmental files? [Y/n]: " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down db port-forwarder&& \
		rm -rf postgres_data && \
		echo "Successful container clean-up"; \
	else \
		echo "Clean up canceled"; \
	fi
migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "Required variable not found. Example: make migrate seq=1 (migration version)"; \
		exit 1; \
	fi; \
	docker compose run --rm migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq "$(seq)"
migrate-up:
	@make migrate-action action=up
migrate-down:
	@make migrate-action action=down
migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "Missing required variable 'action'. Usage: make migrate-up or make migrate-down"; \
		exit 1; \
	fi; \
	docker compose run --rm migrate \
		-path /migrations \
		-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@db:5432/${POSTGRES_DB}?sslmode=disable \
		"$(action)"
env-port-forwarder:
	@docker compose up -d port-forwarder
env-port-down:
	@docker compose down port-forwarder
taskflow-run:
	@export LOGGER_FOLDER=${PROJECT_ROOT}/logs \
	export POSTGRES_HOST=localhost && \
	go mod tidy && \
	go run cmd/*.go