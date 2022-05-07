package models

import (
	"errors"
	"net/http"
)

type HttpError struct {
	Status  int
	Message error
}

func (e HttpError) Error() string {
	return e.Message.Error()
}

func NewBadRequestError(err string) *HttpError {
	return &HttpError{
		Status:  http.StatusBadRequest,
		Message: errors.New(err),
	}
}

func NewUnauthorizedError(err string) *HttpError {
	return &HttpError{
		Status:  http.StatusUnauthorized,
		Message: errors.New(err),
	}
}

func NewInternalServerError(err string) *HttpError {
	return &HttpError{
		Status:  http.StatusInternalServerError,
		Message: errors.New(err),
	}
}
