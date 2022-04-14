package service

import (
	"os"
	"time"
	"todo-app"
	"todo-app/pkg/repository"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

const tokenTTL = 12 * time.Hour

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user todo.User) (int, error) {
	passwordHash, err := s.generatePasswordHash(user.Password)
	if err != nil {
		return 0, err
	}
	user.Password = passwordHash
	return s.repo.CreateUser(user)
}

func (s *AuthService) generatePasswordHash(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (s *AuthService) GenerateToken(email, password string) (string, error) {
	user, err := s.repo.GetUserByEmail(email)
	if err != nil {
		return "", err
	}

	if err := s.compareHashAndPassword(user.Password, password); err != nil {
		return "", err
	}

	token, err := s.generateAccessToken(user.Id)
	return token, err
}

func (s *AuthService) compareHashAndPassword(hashedPassword, password string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)); err != nil {
		return err
	}
	return nil
}

func (s *AuthService) generateAccessToken(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(tokenTTL).Unix(),
		IssuedAt:  time.Now().Unix(),
		Subject:   userId,
	},
	)
	return token.SignedString([]byte(os.Getenv("ACCESS_TOKEN_SECRET")))
}
