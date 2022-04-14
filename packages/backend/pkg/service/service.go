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

type Service struct {
	Authorization
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
	}
}
