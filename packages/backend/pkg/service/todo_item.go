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

func (s *TodoItemService) GetById(userId, itemId int) (models.TodoItem, error) {
	item, err := s.repo.GetById(userId, itemId)
	if err != nil {
		logrus.Errorln("GetById:", err.Error())
		return item, models.NewRequestError(http.StatusBadRequest, errors.New("item not found"))
	}

	return item, nil
}

func (s *TodoItemService) Delete(userId, listId, itemId int) error {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		logrus.Error(err)
		return models.NewRequestError(http.StatusBadRequest, errors.New("list not found"))
	}

	return s.repo.Delete(itemId)
}

func (s *TodoItemService) Update(userId, itemId int, input models.UpdateTodoItemInput) error {
	if input.Done == nil && input.Title == nil {
		return models.NewRequestError(http.StatusBadRequest, errors.New("no input to update"))
	}

	err := s.repo.Update(userId, itemId, input)
	logrus.Error("Update:", err)
	return errors.New("item was not updated")
}
