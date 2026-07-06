include .env
export

export PROJECT_ROOT=$(shell pwd)

env-up:
	docker compose up -d

env-down:
	docker compose down

env-cleanup:
	@read -p "Are you sure you want to cleanup the environment? (y/n): " ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down -v; \
		docker compose rm -f; \
		docker volume rm $(docker volume ls -q); \
		docker network rm $(docker network ls -q); \
	else \
	echo "Environment cleanup cancelled"; \
	fi

run-todo-app:
	@go run cmd/todo_app/main.go