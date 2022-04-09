package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "TodoApp API")
	})

	return router
}
