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

all: test build

build: 
	$(GOBUILD) -o $(BINARY_NAME_SERVER) -v ./cmd/server
	$(GOBUILD) -o $(BINARY_NAME_CLIENT) -v ./cmd/client

test: 
	$(GOTEST) -v ./...

clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME_SERVER)
	rm -f $(BINARY_NAME_CLIENT)

run-server:
	$(GOBUILD) -o $(BINARY_NAME_SERVER) -v ./cmd/server
	./$(BINARY_NAME_SERVER)

run-client:
	$(GOBUILD) -o $(BINARY_NAME_CLIENT) -v ./cmd/client
	./$(BINARY_NAME_CLIENT)

deps:
	$(GOGET) github.com/lib/pq
	$(GOGET) github.com/go-redis/redis/v8
	$(GOGET) github.com/nsf/termbox-go

docker-build:
	docker build -t tank-game-server -f Dockerfile.server .
	docker build -t tank-game-client -f Dockerfile.client .

docker-run:
	$(DOCKER_COMPOSE) up

.PHONY: all build test clean run deps docker-build docker-run