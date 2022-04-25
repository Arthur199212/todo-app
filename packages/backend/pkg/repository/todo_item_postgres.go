package repository

import (
	"fmt"
	"todo-app"

	"github.com/jmoiron/sqlx"
)

type TodoItemPostgres struct {
	db *sqlx.DB
}

func NewTodoItemPostgres(db *sqlx.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

func (r *TodoItemPostgres) Create(userId int, input todo.TodoItemInput) (int, error) {
	query := fmt.Sprintf(`insert into %s (title, list_id)
												values ($1, $2) returning id`, todoItemsTable)
	row := r.db.QueryRow(query, input.Title, input.ListId)
	var id int
	err := row.Scan(&id)
	return id, err
}
