package handler

import (
	"net/http"
	"regexp"
	"todo-app"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func (h *Handler) signUp(c *gin.Context) {
	var input todo.User
	if err := c.BindJSON(&input); err != nil {
		responseWithError(c, todo.NewRequestError(http.StatusBadRequest, err))
		return
	}

	if err := input.Validate(); err != nil {
		responseWithError(c, todo.NewRequestError(http.StatusBadRequest, err))
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		responseWithError(c, todo.NewRequestError(http.StatusInternalServerError, err))
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

type signInInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u signInInput) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.Email, validation.Required, is.Email),
		validation.Field(
			&u.Password,
			validation.Required,
			validation.Length(6, 30),
			validation.Match(regexp.MustCompile("[A-Z]{1}")).Error("should have at least 1 upper case letter"),
			validation.Match(regexp.MustCompile("[0-9]{1}")).Error("should have at least 1 number"),
			validation.Match(regexp.MustCompile("[#?!@$%^&*-]{1}")).Error("should have at least 1 special character"),
		))
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput
	if err := c.BindJSON(&input); err != nil {
		responseWithError(c, todo.NewRequestError(http.StatusBadRequest, err))
		return
	}

	if err := input.Validate(); err != nil {
		responseWithError(c, todo.NewRequestError(http.StatusBadRequest, err))
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Email, input.Password)
	if err != nil {
		responseWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
