package repository

import (
	"fmt"
	"todo-app/models"

	"github.com/jmoiron/sqlx"
)

type TodoListPostgres struct {
	db *sqlx.DB
}

func NewTodoListPostgres(db *sqlx.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r *TodoListPostgres) Create(userId int, title string) (int, error) {
	var id int
	query := fmt.Sprintf(`insert into %s (title, user_id)
												values ($1, $2) returning id`, todoListTable)
	row := r.db.QueryRow(query, title, userId)
	err := row.Scan(&id)
	return id, err
}

func (r *TodoListPostgres) GetAll(userId int) ([]models.TodoList, error) {
	var todoLists []models.TodoList
	query := fmt.Sprintf("select * from %s where user_id=$1", todoListTable)
	err := r.db.Select(&todoLists, query, userId)
	return todoLists, err
}

func (r *TodoListPostgres) GetById(userId, id int) (models.TodoList, error) {
	var todoList models.TodoList
	query := fmt.Sprintf("select * from %s where user_id=$1 and id=$2", todoListTable)
	err := r.db.Get(&todoList, query, userId, id)
	return todoList, err
}

func (r *TodoListPostgres) Update(userId int, input models.UpdateTodoListInput) error {
	query := fmt.Sprintf("update %s set title=$1 where user_id=$2 and id=$3", todoListTable)
	_, err := r.db.Exec(query, input.Title, userId, input.Id)
	return err
}

func (r *TodoListPostgres) Delete(userId, listId int) error {
	query := fmt.Sprintf("delete from %s where user_id=$1 and id=$2", todoListTable)
	// todo: should also delete all related todoItems
	_, err := r.db.Exec(query, userId, listId)
	return err
}
