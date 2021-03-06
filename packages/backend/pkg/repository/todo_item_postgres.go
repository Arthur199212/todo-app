package repository

import (
	"errors"
	"fmt"
	"strings"
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

func (r *TodoItemPostgres) GetById(userId, itemId int) (models.TodoItem, error) {
	var item models.TodoItem
	query := fmt.Sprintf(`select it.id, it.title, it.done, it.list_id from %s it
												inner join %s lt on it.list_id=lt.id and lt.user_id=$1
												where it.id=$2`, todoItemsTable, todoListTable)
	err := r.db.Get(&item, query, userId, itemId)
	return item, err
}

func (r *TodoItemPostgres) Delete(userId, itemId int) error {
	query := fmt.Sprintf(`delete from %s it
												using %s lt 
												where it.id=$1 and lt.user_id=$2 and lt.id=it.list_id`, todoItemsTable, todoListTable)
	res, err := r.db.Exec(query, itemId, userId)
	if err != nil {
		return err
	}

	if rows, err := res.RowsAffected(); err != nil || rows == 0 {
		return errors.New("item was not deleted")
	}

	return err
}

func (r *TodoItemPostgres) Update(userId, itemId int, input models.UpdateTodoItemInput) error {
	setValues := []string{}
	args := []interface{}{}
	argId := 1

	if input.Done != nil {
		setValues = append(setValues, fmt.Sprintf("done=$%d", argId))
		args = append(args, input.Done)
		argId++
	}

	if input.Title != nil {
		setValues = append(setValues, fmt.Sprintf("title=$%d", argId))
		args = append(args, input.Title)
		argId++
	}

	if len(setValues) == 0 {
		return errors.New("no input to update")
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`update %s it set %s
												from %s lt
												where it.id=$%d and lt.id=it.list_id and lt.user_id=$%d`,
		todoItemsTable, setQuery, todoListTable, argId, argId+1)
	args = append(args, itemId, userId)

	res, err := r.db.Exec(query, args...)
	if err != nil {
		return err
	}

	if rows, err := res.RowsAffected(); err != nil || rows == 0 {
		return errors.New("item not found")
	}

	return err
}
