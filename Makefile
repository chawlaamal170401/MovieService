BIN_DIR = bin/
MIGRATE_BIN = $(BIN_DIR)/migrate
SERVER_GO = cmd/server/main.go
CLIENT_GO = cmd/client/main.go
SHELL := /bin/bash

all: migrate build protoc run

.PHONY: migrate
migrate:
	mkdir -p $(BIN_DIR)
	go build -o $(MIGRATE_BIN) ./cmd/migrations
	chmod +x $(MIGRATE_BIN)
	$(MIGRATE_BIN) status | grep 'Pending' > /dev/null && $(MIGRATE_BIN) up || echo "No pending migrations to apply."


.PHONY: build
build:
	@echo "Building server and client..."
	mkdir -p $(BIN_DIR)
	go build -o $(BIN_DIR)/server $(SERVER_GO)
	go build -o $(BIN_DIR)/client $(CLIENT_GO)

.PHONY: clean
clean:
	rm -rf $(BIN_DIR)

.PHONY: run
run: clean migrate build
	@echo "Running server and client..."
	@$(SHELL) -c "(go run $(SERVER_GO) &) && (go run $(CLIENT_GO) &)"

.PHONY: run_server
run_server:
	go run $(SERVER_GO)

.PHONY: run_client
run_client:
	go run $(CLIENT_GO)

.PHONY: stop
stop:
	@echo "Stopping server and client..."
	@SERVER_PID=$(shell ps aux | grep 'go run $(SERVER_GO)' | grep -v 'grep' | awk '{print $$2}') && \
		if [ ! -z "$$SERVER_PID" ]; then kill -9 $$SERVER_PID; fi || echo "Server not running."
	@CLIENT_PID=$(shell ps aux | grep 'go run $(CLIENT_GO)' | grep -v 'grep' | awk '{print $$2}') && \
		if [ ! -z "$$CLIENT_PID" ]; then kill -9 $$CLIENT_PID; fi || echo "Client not running."
	@sleep 1  # Give time for processes to terminate


.PHONY: protoc
protoc:
	@protoc --proto_path=internals/proto \
		--go_out=internals/proto --go-grpc_out=internals/proto \
		--go_opt=paths=source_relative --go-grpc_opt=paths=source_relative \
		--grpc-gateway_out=internals/proto --grpc-gateway_opt=logtostderr=true --grpc-gateway_opt=paths=source_relative \
		internals/proto/movie.proto
