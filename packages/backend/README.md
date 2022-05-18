# Backend

### Technologies

1. Golang
1. Gin
1. JWT
1. bcrypt
1. ozzo-validation
1. SQL

### How to start

1. If you start app for the first time, create docker volume
   `docker create volume pgdata`
1. Spin up PostgreSQL DB in docker container
   `make up`
1. If you start app for the first time, run migrations
   `make migrate`
1. Start the server
   `make run`

### How to run migrations with [migrate](https://github.com/golang-migrate/migrate)

```bash
# Create migration
# -ext -- extension of the file
# -dir -- directory
# -seq -- migration name
migrate create -ext sql -dir ./schema -seq init

# Apply migration
migrate -path ./schema -database 'postgres://postgres:secret@localhost:7010/postgres?sslmode=disable' up

# Rollback migration
migrate -path ./schema -database 'postgres://postgres:secret@localhost:7010/postgres?sslmode=disable' down
```
