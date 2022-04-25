package handler

import (
	"net/http"
	"todo-app/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllItems(c *gin.Context) {}

func (h *Handler) createItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		responseWithError(c, models.NewRequestError(http.StatusUnauthorized, err))
		return
	}

	var input models.TodoItemInput
	if err := c.BindJSON(&input); err != nil {
		responseWithError(c, models.NewRequestError(http.StatusBadRequest, err))
		return
	}

	if err := input.Validate(); err != nil {
		responseWithError(c, models.NewRequestError(http.StatusBadRequest, err))
		return
	}

	id, err := h.services.TodoItem.Create(userId, input)
	if err != nil {
		responseWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, id)
}

func (h *Handler) getItemById(c *gin.Context) {}

func (h *Handler) updateItem(c *gin.Context) {}

func (h *Handler) deleteItem(c *gin.Context) {}
