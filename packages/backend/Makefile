run:
	go run ./cmd/main.go

build:
	npx rimraf bin && go build -v -o bin/main ./cmd

test:
	go test -v ./...

migrate:
	migrate -path ./schema -database 'postgres://postgres:secret@localhost:7010/postgres?sslmode=disable' up

.DEFAULT_GOAL := build
