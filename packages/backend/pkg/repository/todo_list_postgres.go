package repository

import (
	"fmt"
	"todo-app"

	"github.com/jmoiron/sqlx"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r *TodoListPostgres) Create(userId int, todoList todo.CreateListInput) (int, error) {
	var id int
	query := fmt.Sprintf(`insert into %s (title, user_id)
												values ($1, $2) returning id`, todoListTable)
	row := r.db.QueryRow(query, todoList.Title, todoList.UserId)
	err := row.Scan(&id)
	return id, err
}

func (r *TodoListPostgres) GetAll(userId int) ([]todo.TodoList, error) {
	var todoLists []todo.TodoList
	query := fmt.Sprintf("select * from %s where user_id=$1", todoListTable)
	err := r.db.Select(&todoLists, query, userId)
	return todoLists, err
}
