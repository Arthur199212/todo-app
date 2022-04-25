package service

import (
	"errors"
	"net/http"
	"todo-app"
	"todo-app/pkg/repository"

	"github.com/sirupsen/logrus"
)

type TodoItemService struct {
	repo     repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItem(repo repository.TodoItem, listRepo repository.TodoList) *TodoItemService {
	return &TodoItemService{repo: repo, listRepo: listRepo}
}

func (s *TodoItemService) Create(userId int, input todo.TodoItemInput) (int, error) {
	_, err := s.listRepo.GetById(userId, input.ListId)
	if err != nil {
		logrus.Error(err)
		return 0, todo.NewRequestError(http.StatusBadRequest, errors.New("list not found"))
	}

	return s.repo.Create(userId, input)
}
