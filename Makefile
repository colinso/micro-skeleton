DB_HOST?=localhost
DB_USER?=user
DB_PASSWORD?=password
DB_NAME?=skeleton

GOOSE_DB_STRING?="host=${DB_HOST} user=${DB_USER} password=${DB_PASSWORD} dbname=${DB_NAME} sslmode=disable"

# BUILD AND START SERVICES

build:
	go build -o bin/skeleton cmd/main.go

run: generate build
	export APP_ENV=dev && ./bin/skeleton

up-deps:
	docker-compose up -d

down-deps:
	docker-compose down

generate:
	go generate ./...

# DATABASE

migrate:
	goose -dir ./internal/db/migrations postgres ${GOOSE_DB_STRING} up

migrate-down:
	goose -dir ./internal/db/migrations postgres ${GOOSE_DB_STRING} down

migration-status:
	goose -dir ./internal/db/migrations postgres ${GOOSE_DB_STRING} status

migration-create:
	goose -dir ./internal/db/migrations postgres ${GOOSE_DB_STRING} create temp_name sql

db-connect:
	psql -h ${DB_HOST} -U ${DB_USER} ${DB_NAME}