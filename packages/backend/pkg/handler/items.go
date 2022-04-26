package handler

import (
	"errors"
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
		responseWithError(c, models.NewRequestError(http.StatusBadRequest, errors.New("listId is invalid")))
		return
	}

	err = validation.Validate(listId,
		validation.Required.Error("listId is required"),
		validation.Min(0).Error("listId is invalid"),
	)
	if err != nil {
		responseWithError(c, models.NewRequestError(http.StatusBadRequest, err))
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
		responseWithError(c, models.NewRequestError(http.StatusBadRequest, errors.New("listId is invalid")))
		return
	}

	var input models.TodoItemInput
	if err := c.BindJSON(&input); err != nil {
		responseWithError(c, models.NewRequestError(http.StatusBadRequest, err))
		return
	}
	input.ListId = listId

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

func (h *Handler) getItemById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		responseWithError(c, err)
		return
	}

	itemId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responseWithError(c, models.NewRequestError(http.StatusBadRequest, err))
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
		responseWithError(c, models.NewRequestError(http.StatusBadRequest, err))
		return
	}

	var input models.UpdateTodoItemInput
	if err := c.BindJSON(&input); err != nil {
		responseWithError(c, models.NewRequestError(http.StatusBadRequest, err))
		return
	}

	if err := validation.Validate(itemId,
		validation.Required.Error("itemId is required"),
		validation.Min(0).Error("itemId is required"),
	); err != nil {
		responseWithError(c, models.NewRequestError(http.StatusBadRequest, err))
		return
	}

	if err := input.Validate(); err != nil {
		responseWithError(c, models.NewRequestError(http.StatusBadRequest, err))
		return
	}

	if err := h.services.TodoItem.Update(userId, itemId, input); err != nil {
		responseWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"message": "OK"})
}

type deleteItemInput struct {
	ItemId int `json:"itemId"`
	ListId int `json:"listId"`
}

func (input deleteItemInput) Validate() error {
	return validation.ValidateStruct(&input,
		validation.Field(&input.ItemId, validation.Required, validation.Min(0)),
		validation.Field(&input.ItemId, validation.Required, validation.Min(0)),
	)
}

func (h *Handler) deleteItem(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		responseWithError(c, err)
		return
	}

	var input deleteItemInput
	if err := c.BindJSON(&input); err != nil {
		responseWithError(c, models.NewRequestError(http.StatusBadRequest, err))
		return
	}

	if err := input.Validate(); err != nil {
		responseWithError(c, models.NewRequestError(http.StatusBadRequest, err))
		return
	}

	if err := h.services.TodoItem.Delete(userId, input.ListId, input.ItemId); err != nil {
		responseWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"message": "OK"})
}
