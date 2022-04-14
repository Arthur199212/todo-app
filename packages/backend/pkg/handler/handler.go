package handler

import (
	"net/http"
	"todo-app/pkg/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "TodoApp API")
	})

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	api := router.Group("/api", h.authRequired)
	{
		lists := api.Group("/lists")
		{
			lists.GET("/", h.getAllLists)
			lists.POST("/", h.createList)
			lists.GET("/:id", h.getListById)
			lists.PUT("/:id", h.updateList)
			lists.DELETE("/:id", h.deleteList)
		}

		items := api.Group("/items")
		{
			items.GET("/", h.getAllItems)
			items.POST("/", h.createItem)
			items.GET("/:id", h.getItemById)
			items.PUT("/:id", h.updateItem)
			items.DELETE("/:id", h.deleteItem)
		}
	}

	return router
}
