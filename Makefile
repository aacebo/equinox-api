include .env.$(GO_ENV)
export

POSTGRES_CONNECTION_STRING := postgres://$(PG_USER):$(PG_PASSWORD)@$(PG_HOST):5432/$(PG_DATABASE)?sslmode=disable

build:
	go build -o bin/main src/main.go

clean:
	rm -rf ./bin

run:
	go run src/main.go

migrate-up:
	migrate -source file://migrations -database $(POSTGRES_CONNECTION_STRING) up

migrate-down:
	migrate -source file://migrations -database $(POSTGRES_CONNECTION_STRING) down

migrate-new:
	migrate create -ext sql -dir migrations $(name)
