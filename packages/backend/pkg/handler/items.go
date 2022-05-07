package handler

import (
	"net/http"
	"strconv"
	"todo-app/models"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (h *Handler) getAllItems(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		responseWithError(c, err)
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responseWithError(c, models.NewBadRequestError("listId is invalid"))
		return
	}

	err = validation.Validate(listId,
		validation.Required.Error("listId is required"),
		validation.Min(0).Error("listId is invalid"),
	)
	if err != nil {
		responseWithError(c, models.NewBadRequestError(err.Error()))
		return
	}

	items, err := h.services.TodoItem.GetAllByListId(userId, listId)
	if err != nil {
		responseWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, items)
}

func (h *Handler) createItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		responseWithError(c, err)
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responseWithError(c, models.NewBadRequestError("listId is invalid"))
		return
	}

	var input models.TodoItemInput
	if err := c.BindJSON(&input); err != nil {
		responseWithError(c, models.NewBadRequestError(err.Error()))
		return
	}
	input.ListId = listId

	if err := input.Validate(); err != nil {
		responseWithError(c, models.NewBadRequestError(err.Error()))
		return
	}

	id, err := h.services.TodoItem.Create(userId, input)
	if err != nil {
		responseWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"id": id})
}

func (h *Handler) getItemById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		responseWithError(c, err)
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responseWithError(c, models.NewBadRequestError(err.Error()))
		return
	}

	item, err := h.services.TodoItem.GetById(userId, itemId)
	if err != nil {
		responseWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, item)
}

func (h *Handler) updateItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		responseWithError(c, err)
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responseWithError(c, models.NewBadRequestError(err.Error()))
		return
	}

	var input models.UpdateTodoItemInput
	if err := c.BindJSON(&input); err != nil {
		responseWithError(c, models.NewBadRequestError(err.Error()))
		return
	}

	if err := input.Validate(); err != nil {
		responseWithError(c, models.NewBadRequestError(err.Error()))
		return
	}

	if err := h.services.TodoItem.Update(userId, itemId, input); err != nil {
		responseWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"message": "OK"})
}

func (h *Handler) deleteItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		responseWithError(c, err)
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responseWithError(c, models.NewBadRequestError(err.Error()))
		return
	}

	if err := h.services.TodoItem.Delete(userId, itemId); err != nil {
		responseWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"message": "OK"})
}
