package service

import (
	"errors"
	"net/http"
	"os"
	"strconv"
	"time"
	"todo-app/models"
	"todo-app/pkg/repository"

	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

const tokenTTL = 12 * time.Hour

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.User) (int, error) {
	passwordHash, err := s.generatePasswordHash(user.Password)
	if err != nil {
		logrus.Errorln("CreateUser:", err)
		return 0, models.NewRequestError(http.StatusInternalServerError, errors.New("password is not valid"))
	}
	user.Password = passwordHash

	id, err := s.repo.CreateUser(user)
	if err != nil {
		return 0, models.NewRequestError(http.StatusInternalServerError, errors.New("could not create user"))
	}
	return id, nil
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
		logrus.Errorln("GenerateToken:", err)
		return "", models.NewRequestError(http.StatusBadRequest, errors.New("email or password are not correct"))
	}

	if err := s.compareHashAndPassword(user.Password, password); err != nil {
		logrus.Errorln("GenerateToken:", err)
		return "", models.NewRequestError(http.StatusBadRequest, errors.New("email or password are not correct"))
	}

	token, err := s.generateAccessToken(strconv.Itoa(user.Id))
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

func (s *AuthService) ParseUserIdFromToken(tokenStr string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid jwt token signing method")
		}
		return []byte(os.Getenv("ACCESS_TOKEN_SECRET")), nil
	})
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*jwt.StandardClaims); ok && token.Valid {
		return strconv.Atoi(claims.Subject)
	}
	return 0, errors.New("token is invalid")
}
