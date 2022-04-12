package service

import (
	"todo-app"
	"todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(input todo.User) (int, error)
}
type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
