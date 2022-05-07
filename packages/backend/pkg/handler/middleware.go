package handler

import (
	"strings"
	"todo-app/models"

	"github.com/gin-gonic/gin"
)

const (
	authHeader = "Authorization"
	userCtx    = "userId"
)

func (h *Handler) authRequired(c *gin.Context) {
	header := strings.Split(c.GetHeader(authHeader), " ")
	if len(header) != 2 {
		responseWithError(c, models.NewBadRequestError("not authorized"))
		return
	}
	token := header[1]
	userId, err := h.services.Authorization.ParseUserIdFromToken(token)
	if err != nil {
		responseWithError(c, models.NewBadRequestError(err.Error()))
		return
	}

	c.Set(userCtx, userId)
}

func getUserId(c *gin.Context) (int, error) {
	userId, ok := c.Get(userCtx)
	if !ok {
		return 0, models.NewUnauthorizedError("userId is invalid")
	}

	userIdInt, ok := userId.(int)
	if !ok {
		return 0, models.NewUnauthorizedError("userId is invalid")
	}

	return userIdInt, nil
}
