package service

import "todo-app/pkg/repository"

type AuthService struct {
	repo *repository.Authorization
}

func NewAuthService(repo *repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}
