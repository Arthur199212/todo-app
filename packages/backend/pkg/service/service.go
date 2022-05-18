package service

import (
	"todo-app/models"
	"todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(input models.User) (int, error)
	GenerateToken(email, password string) (string, error)
	ParseUserIdFromToken(token string) (int, error)
}

type TodoList interface {
	Create(userId int, title string) (int, error)
	GetAll(userId int) ([]models.TodoList, error)
	GetById(userId, listId int) (models.TodoList, error)
	Update(userId int, input models.UpdateTodoListInput) error
	Delete(userId, listId int) error
}

type TodoItem interface {
	Create(userId int, input models.TodoItemInput) (int, error)
	GetAllByListId(userId, listId int) ([]models.TodoItem, error)
	GetById(userId, itemId int) (models.TodoItem, error)
	Delete(userId, itemId int) error
	Update(userId, itemId int, input models.UpdateTodoItemInput) error
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
