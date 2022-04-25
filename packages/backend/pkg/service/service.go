package service

import (
	"todo-app"
	"todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(input todo.User) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseUserIdFromToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, title string) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId, listId int) (todo.TodoList, error)
	Update(userId int, input todo.UpdateTodoListInput) error
	Delete(userId, listId int) error
}

type TodoItem interface {
	Create(userId int, input todo.TodoItemInput) (int, error)
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoList(repos.TodoList, repos.Authorization),
		TodoItem:      NewTodoItem(repos.TodoItem, repos.TodoList),
	}
}
