package repository

import (
	"todo-app"

	"github.com/jmoiron/sqlx"
)

const (
	usersTable     = "users"
	todoListTable  = "todo_lists"
	todoItemsTable = "todo_items"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GetUserByEmail(email string) (todo.User, error)
}

type Repository struct {
	Authorization
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
