package repository

import (
	"fmt"
	"todo-app/models"

	"github.com/jmoiron/sqlx"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

func (r *TodoItemPostgres) Create(userId int, input models.TodoItemInput) (int, error) {
	var id int
	query := fmt.Sprintf(`insert into %s (title, list_id)
												values ($1, $2) returning id`, todoItemsTable)
	row := r.db.QueryRow(query, input.Title, input.ListId)
	err := row.Scan(&id)
	return id, err
}

func (r *TodoItemPostgres) GetAllByListId(listId int) ([]models.TodoItem, error) {
	var items []models.TodoItem
	query := fmt.Sprintf("select * from %s where list_id=$1", todoItemsTable)
	err := r.db.Select(&items, query, listId)
	return items, err
}
