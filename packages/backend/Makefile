run: up
	go run ./cmd/main.go

clean:
	npx rimraf bin cover.out

build: clean
	go build -v -o bin/main.exe ./cmd

test:
	go test --short -coverprofile=cover.out -v ./...

up:
	docker compose up -d

down:
	docker compose down

migrate:
	migrate -path ./schema -database 'postgres://postgres:secret@localhost:7010/postgres?sslmode=disable' up

gen:
	mockgen -source=pkg/service/service.go -destination=pkg/service/mocks/mock.go

.DEFAULT_GOAL := run
