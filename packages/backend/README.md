# Backend

### How to start

1. Create docker volume
`docker create volume pgdata`
2. Spin up PostgreSQL DB in docker container
`npm run up`
3. Build the code
`npm run build`
4. Start the server
`./bin/main.exe`

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
