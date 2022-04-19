package service

import (
	"errors"
	"todo-app"
	"todo-app/pkg/repository"
)

type TodoListService struct {
	repo     repository.TodoList
	authRepo repository.Authorization
}

func NewTodoList(repo repository.TodoList, authRepo repository.Authorization) *TodoListService {
	return &TodoListService{repo: repo, authRepo: authRepo}
}

func (s *TodoListService) Create(userId int, todoList todo.CreateListInput) (int, error) {
	if _, err := s.authRepo.GetUserById(userId); err != nil {
		return 0, errors.New("user does not exists")
	}
	return s.repo.Create(userId, todoList)
}

func (s *TodoListService) GetAll(userId int) ([]todo.TodoList, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetById(userId, id int) (todo.TodoList, error) {
	return s.repo.GetById(userId, id)
}
