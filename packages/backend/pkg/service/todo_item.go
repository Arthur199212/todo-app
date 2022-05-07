package service

import (
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
		logrus.Errorln("Create:", err)
		return 0, models.NewBadRequestError("list not found")
	}

	return s.repo.Create(userId, input)
}

func (s *TodoItemService) GetAllByListId(userId, listId int) ([]models.TodoItem, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		logrus.Errorln("GetAllByListId:", err)
		return nil, models.NewBadRequestError("list not found")
	}

	return s.repo.GetAllByListId(listId)
}

func (s *TodoItemService) GetById(userId, itemId int) (models.TodoItem, error) {
	item, err := s.repo.GetById(userId, itemId)
	if err != nil {
		logrus.Errorln("GetById:", err.Error())
		return item, models.NewBadRequestError("item not found")
	}

	return item, nil
}

func (s *TodoItemService) Delete(userId, itemId int) error {
	if err := s.repo.Delete(userId, itemId); err != nil {
		logrus.Errorln("Delete:", err)
		return models.NewBadRequestError("item was not deleted")
	}

	return nil
}

func (s *TodoItemService) Update(userId, itemId int, input models.UpdateTodoItemInput) error {
	if input.Done == nil && input.Title == nil {
		return models.NewBadRequestError("no input to update")
	}

	err := s.repo.Update(userId, itemId, input)
	if err != nil {
		logrus.Errorln("Update:", err)
		return models.NewBadRequestError("item was not updated")
	}

	return nil
}
