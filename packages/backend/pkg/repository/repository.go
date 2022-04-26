package repository

import (
	"todo-app/models"

	"github.com/jmoiron/sqlx"
)

const (
	usersTable     = "users"
	todoListTable  = "todo_lists"
	todoItemsTable = "todo_items"
)

type Authorization interface {
	CreateUser(user models.User) (int, error)
	GetUserByEmail(email string) (models.User, error)
	GetUserById(id int) (models.User, error)
}

type TodoList interface {
	Create(userId int, title string) (int, error)
	GetAll(userId int) ([]models.TodoList, error)
	GetById(userId, id int) (models.TodoList, error)
	Update(userId int, input models.UpdateTodoListInput) error
	Delete(userId, listId int) error
}

type TodoItem interface {
	Create(userId int, input models.TodoItemInput) (int, error)
	GetAllByListId(listId int) ([]models.TodoItem, error)
	GetById(userId, itemId int) (models.TodoItem, error)
	Delete(id int) error
	Update(itemId int, input models.UpdateTodoItemInput) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
