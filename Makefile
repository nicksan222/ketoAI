.PHONY: help build dev kill generate new-model

help: ## Display this help message
	@echo "Usage: make [TARGET]\n"
	@echo "Targets:"
	@echo "  build       Build the Docker containers."
	@echo "  dev         Start the Docker containers in the background."
	@echo "  test-backend        Run the tests. in backend"
	@echo "  test-frontend        Run the tests. in frontend"
	@echo "  kill        Stop and remove the Docker containers."
	@echo "  generate    Run go generate command inside the app container."
	@echo "  new-model   Create a new model. Usage: make new-model model=ModelName"
	@echo "  schema      Generate the database schema diagram."
	@echo "  db-shell    Open a shell to the database."

build: ## Build the Docker containers
	@docker compose build

dev: ## Start the Docker containers in the background
	@docker compose up -d --remove-orphans

test-backend:
	@docker compose run --rm backend go test ./...

test-backend-coverage:
	@docker compose run --rm backend go test ./... -cover

kill: ## Stop and remove the Docker containers
	@docker compose down

db-shell:
	@docker-compose exec postgres psql -U root -d ketoai

seed-shell:
	@docker-compose exec seed bash

logs-backend: ## Show and follow the logs for the backend container
	@docker-compose logs -f backend

logs-frontend: ## Show and follow the logs for the frontend container
	@docker-compose logs -f frontend

logs-seed: ## Show and follow the logs for the seed container
	@docker-compose logs -f seed

logs-mongodb: ## Show and follow the logs for the mongodb container
	@docker-compose logs -f mongodb

frontend-shell:
	@docker-compose exec frontend /bin/sh

gen-api-client:
	cd backend
	@make all
	cd ..
	cd frontend
	@make generate-api
	cd ..