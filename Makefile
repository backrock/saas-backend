.PHONY: help build run test clean docker-up docker-down migrate

help:
	@echo "Available commands:"
	@echo "  make build           - Build the application"
	@echo "  make run             - Run the application"
	@echo "  make test            - Run tests"
	@echo "  make clean           - Clean build artifacts"
	@echo "  make docker-up       - Start Docker containers"
	@echo "  make docker-down     - Stop Docker containers"
	@echo "  make docker-build    - Build Docker image"
	@echo "  make migrate         - Run database migrations"
	@echo "  make sqlc            - Generate code from SQL"
	@echo "  make seed            - Seed database with initial data"

build:
	go build -o bin/server ./cmd/server

run:
	go run ./cmd/server

test:
	go test -v -cover ./...

clean:
	rm -rf bin/
	go clean

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

docker-build:
	docker-compose build

docker-logs:
	docker-compose logs -f app

migrate:
	@echo "Running database migrations..."
	@# TODO: Implement migration runner

sqlc:
	sqlc generate

seed:
	@echo "Seeding database..."
	@# TODO: Implement database seeding

fmt:
	go fmt ./...

lint:
	golangci-lint run

vet:
	go vet ./...

mod-tidy:
	go mod tidy

deps:
	go mod download
	go mod verify
