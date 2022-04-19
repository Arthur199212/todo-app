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
	GetUserById(id int) (todo.User, error)
}

type TodoList interface {
	Create(userId int, todoList todo.CreateListInput) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
}

type Repository struct {
	Authorization
	TodoList
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList: NewTodoListPostgres(db),
	}
}
