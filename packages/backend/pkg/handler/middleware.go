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
		responseWithError(c, models.NewUnauthorizedError("not authorized"))
		return
	}

	if header[0] != "Bearer" || header[1] == "" {
		responseWithError(c, models.NewUnauthorizedError("invalid auth header"))
		return
	}

	token := header[1]
	userId, err := h.services.Authorization.ParseUserIdFromToken(token)
	if err != nil {
		responseWithError(c, models.NewUnauthorizedError("invalid token"))
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
