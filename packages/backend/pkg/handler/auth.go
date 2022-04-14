package handler

import (
	"net/http"
	"regexp"
	"todo-app"

	"github.com/gin-gonic/gin"
	v "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
)

func (h *Handler) signUp(c *gin.Context) {
	var input todo.User
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := input.Validate(); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
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
	return v.ValidateStruct(&u,
		v.Field(&u.Email, v.Required, is.Email),
		v.Field(&u.Password, v.Required, v.Length(6, 30)),
		v.Field(&u.Password, v.Required, v.Match(regexp.MustCompile("[A-Z]{1}")).Error("should have at least 1 upper case letter")),
		v.Field(&u.Password, v.Required, v.Match(regexp.MustCompile("[0-9]{1}")).Error("should have at least 1 number")),
		v.Field(&u.Password, v.Required, v.Match(regexp.MustCompile("[#?!@$%^&*-]{1}")).Error("should have at least 1 special character")),
	)
}

func (h *Handler) signIn(c *gin.Context) {
	var input signInInput
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	if err := input.Validate(); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.services.Authorization.GenerateToken(input.Email, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})
}
