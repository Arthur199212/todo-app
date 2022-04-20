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

func (s *TodoListService) Create(userId int, title string) (int, error) {
	if _, err := s.authRepo.GetUserById(userId); err != nil {
		return 0, errors.New("user does not exists")
	}
	return s.repo.Create(userId, title)
}

func (s *TodoListService) GetAll(userId int) ([]todo.TodoList, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetById(userId, listId int) (todo.TodoList, error) {
	return s.repo.GetById(userId, listId)
}

func (s *TodoListService) Update(userId int, input todo.UpdateTodoListInput) error {
	return s.repo.Update(userId, input)
}

func (s *TodoListService) Delete(userId, listId int) error {
	return s.repo.Delete(userId, listId)
}
