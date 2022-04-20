package todo

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type TodoList struct {
	Id     int    `json:"id" db:"id"`
	Title  string `json:"title" db:"title"`
	UserId int    `json:"userId" db:"user_id"`
}

func (tl TodoList) Validate() error {
	return validation.ValidateStruct(&tl,
		validation.Field(&tl.Id, validation.Required, validation.Min(0)),
		validation.Field(&tl.Title, validation.Required, validation.Length(3, 50)),
		validation.Field(&tl.UserId, validation.Required, validation.Min(0)),
	)
}

type UpdateTodoListInput struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

func (input UpdateTodoListInput) Validate() error {
	return validation.ValidateStruct(&input,
		validation.Field(&input.Id, validation.Required, validation.Min(0)),
		validation.Field(&input.Title, validation.Required, validation.Length(3, 50)),
	)
}

type TodoItem struct {
	Id     int    `json:"id" db:"id"`
	Title  string `json:"title" db:"title"`
	Done   bool   `json:"done" db:"done"`
	ListId int    `json:"listId" db:"list_id"`
}
