BIN_DIR = bin/
MIGRATE_BIN = $(BIN_DIR)/migrate
SERVER_GO = cmd/server/main.go
CLIENT_GO = cmd/client/main.go
SHELL := /bin/bash

all: build up run_server run_client protoc

.PHONY: build
build:
	mkdir -p $(BIN_DIR)
	go build -o $(MIGRATE_BIN) ./cmd/migrations
	chmod +x $(MIGRATE_BIN)

.PHONY: up
up: build
	$(MIGRATE_BIN) up

.PHONY: clean
clean:
	rm -rf $(BIN_DIR)

.PHONY: run_server
run_server:
	go run $(SERVER_GO)

.PHONY: run_client
run_client:
	go run $(CLIENT_GO)

.PHONY: protoc
protoc:
	@protoc --proto_path=internals/proto \
		--go_out=internals/proto --go-grpc_out=internals/proto \
		--go_opt=paths=source_relative --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=internals/proto --grpc-gateway_opt=logtostderr=true --grpc-gateway_opt=paths=source_relative \
		internals/proto/movie.proto
