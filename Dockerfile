# Use an official Go runtime as a parent image
FROM golang:1.23 AS builder


RUN apt-get update && apt-get install -y \
    postgresql postgresql-contrib \
    && apt-get clean

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum are not changed
RUN go mod tidy

# Copy the entire source code into the container
COPY . .

# Set environment variables
ENV GOPATH=/go
ENV PATH=$PATH:/go/bin

# Install Protobuf compiler and Go plugins
RUN apt-get update && apt-get install -y \
    protobuf-compiler \
    && go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28.0 \
    && go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2.0 \
    && go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@v2.9.0

# Build the migrate binary
RUN mkdir -p bin && go build -o bin/migrate ./cmd/migrations && chmod +x bin/migrate

# Build the server and client
RUN go build -o bin/server ./cmd/server/main.go && \
    go build -o bin/client ./cmd/client/main.go

# Run migrations before starting the server
RUN bin/migrate status | grep 'Pending' > /dev/null && bin/migrate up || echo "No pending migrations to apply."

# Expose the port the app runs on
EXPOSE 8000

# Copy the start script
COPY start.sh /app/start.sh
RUN chmod +x /app/start.sh

# Run the start script
CMD ["/app/start.sh"]