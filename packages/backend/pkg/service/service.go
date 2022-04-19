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
	Create(userId int, todoList todo.CreateListInput) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
}

type Service struct {
	Authorization
	TodoList
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList: NewTodoList(repos.TodoList, repos.Authorization),
	}
}
