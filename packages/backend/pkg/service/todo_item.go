package service

import (
	"errors"
	"net/http"
	"todo-app/models"
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

func (s *TodoItemService) Create(userId int, input models.TodoItemInput) (int, error) {
	_, err := s.listRepo.GetById(userId, input.ListId)
	if err != nil {
		logrus.Error(err)
		return 0, models.NewRequestError(http.StatusBadRequest, errors.New("list not found"))
	}

	return s.repo.Create(userId, input)
}

func (s *TodoItemService) GetAllByListId(userId, listId int) ([]models.TodoItem, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		logrus.Error(err)
		return nil, models.NewRequestError(http.StatusBadRequest, errors.New("list not found"))
	}

	return s.repo.GetAllByListId(listId)
}

func (s *TodoItemService) GetById(userId, listId, itemId int) (models.TodoItem, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		logrus.Error(err)
		return models.TodoItem{}, models.NewRequestError(http.StatusBadRequest, errors.New("list not found"))
	}

	return s.repo.GetById(itemId)
}
