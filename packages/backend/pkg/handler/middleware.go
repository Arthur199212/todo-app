package handler

import (
	"errors"
	"net/http"
	"strings"
	"todo-app"

	"github.com/gin-gonic/gin"
)

const (
	authHeader = "Authorization"
	userCtx    = "userId"
)

func (h *Handler) authRequired(c *gin.Context) {
	header := strings.Split(c.GetHeader(authHeader), " ")
	if len(header) != 2 {
		responseWithError(c, todo.NewRequestError(http.StatusBadRequest, errors.New("not authorized")))
		return
	}
	token := header[1]
	userId, err := h.services.Authorization.ParseUserIdFromToken(token)
	if err != nil {
		responseWithError(c, todo.NewRequestError(http.StatusBadRequest, err))
		return
	}

	c.Set(userCtx, userId)
}

func getUserId(c *gin.Context) (int, error) {
	userId, ok := c.Get(userCtx)
	if !ok {
		return 0, errors.New("userId is invalid")
	}

	userIdInt, ok := userId.(int)
	if !ok {
		return 0, errors.New("userId is invalid")
	}

	return userIdInt, nil
}
