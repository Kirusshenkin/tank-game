# Makefile

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME_SERVER=tank-server
BINARY_NAME_CLIENT=tank-client
DOCKER_COMPOSE=docker-compose
BUILD_DIR=build

# Database connection parameters
DB_USER=postgres
DB_NAME=postgres
DB_PASSWORD=postgres
DB_HOST=localhost
DB_PORT=5432
DB_SSLMODE=disable

.PHONY: all build clean test run run-server run-client deps docker-build docker-run docker-stop docker-restart docker-logs check rebuild-and-restart init-db migrate-up migrate-down migrate-status migrate-create

all: test build

# Build the binaries and place them in the build directory
build:
	mkdir -p $(BUILD_DIR)
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME_SERVER) ./cmd/server
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME_CLIENT) ./cmd/client

clean:
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)

test:
	$(GOTEST) -v ./...

run-server:
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME_SERVER) ./cmd/server
	chmod +x $(BUILD_DIR)/$(BINARY_NAME_SERVER)
	./$(BUILD_DIR)/$(BINARY_NAME_SERVER)

run-client:
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME_CLIENT) ./cmd/client
	./$(BUILD_DIR)/$(BINARY_NAME_CLIENT)

deps:
	$(GOCMD) mod download

docker-build:
	$(DOCKER_COMPOSE) build

docker-run:
	$(DOCKER_COMPOSE) up -d

docker-stop:
	$(DOCKER_COMPOSE) down

docker-restart: docker-stop docker-run

docker-logs:
	$(DOCKER_COMPOSE) logs -f

check: test
	go vet ./...
	golint ./...

rebuild-and-restart: clean build docker-build docker-restart docker-logs

init-db:
	docker exec -it $$(docker ps -qf "name=postgres") psql -U $(DB_USER) -c "CREATE USER tankuser WITH PASSWORD 'tankpassword';"
	docker exec -it $$(docker ps -qf "name=postgres") psql -U $(DB_USER) -c "CREATE DATABASE tankgame;"
	docker exec -it $$(docker ps -qf "name=postgres") psql -U $(DB_USER) -c "GRANT ALL PRIVILEGES ON DATABASE tankgame TO tankuser;"
	docker exec -it $$(docker ps -qf "name=postgres") psql -U $(DB_USER) -d tankgame -c "GRANT ALL PRIVILEGES ON ALL TABLES IN SCHEMA public TO tankuser;"
	docker exec -it $$(docker ps -qf "name=postgres") psql -U $(DB_USER) -d tankgame -c "GRANT ALL PRIVILEGES ON ALL SEQUENCES IN SCHEMA public TO tankuser;"

migrate-up:
	goose -dir ./internal/storage/migrations postgres "user=$(DB_USER) password=$(DB_PASSWORD) dbname=$(DB_NAME) host=$(DB_HOST) port=$(DB_PORT) sslmode=$(DB_SSLMODE)" up

migrate-down:
	goose -dir ./internal/storage/migrations postgres "user=$(DB_USER) password=$(DB_PASSWORD) dbname=$(DB_NAME) host=$(DB_HOST) port=$(DB_PORT) sslmode=$(DB_SSLMODE)" down

migrate-status:
	goose -dir ./internal/storage/migrations postgres "user=$(DB_USER) password=$(DB_PASSWORD) dbname=$(DB_NAME) host=$(DB_HOST) port=$(DB_PORT) sslmode=$(DB_SSLMODE)" status

migrate-create:
	@read -p "Enter migration name: " name; \
	goose -dir ./internal/storage/migrations create $$name sql

rebuild-and-restart: docker-stop docker-build docker-run init-db
