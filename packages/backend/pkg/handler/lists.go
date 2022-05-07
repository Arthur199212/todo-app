package handler

import (
	"net/http"
	"strconv"
	"todo-app/models"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (h *Handler) getAllLists(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		responseWithError(c, err)
		return
	}

	// todo: pagination

	todoLists, err := h.services.TodoList.GetAll(userId)
	if err != nil {
		responseWithError(c, models.NewInternalServerError(err.Error()))
	}

	c.JSON(http.StatusOK, todoLists)
}

type todoListInput struct {
	Title string `json:"title"`
}

func (input todoListInput) Validate() error {
	return validation.ValidateStruct(&input,
		validation.Field(&input.Title, validation.Required, validation.Length(3, 50)),
	)
}

func (h *Handler) createList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		responseWithError(c, err)
		return
	}

	var input todoListInput
	if err := c.BindJSON(&input); err != nil {
		responseWithError(c, models.NewBadRequestError(err.Error()))
		return
	}

	if err := input.Validate(); err != nil {
		responseWithError(c, models.NewBadRequestError(err.Error()))
		return
	}

	listId, err := h.services.TodoList.Create(userId, input.Title)
	if err != nil {
		responseWithError(c, models.NewInternalServerError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"id": listId})
}

func (h *Handler) getListById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		responseWithError(c, err)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responseWithError(c, models.NewBadRequestError(err.Error()))
		return
	}

	todoList, err := h.services.TodoList.GetById(userId, id)
	if err != nil {
		responseWithError(c, models.NewBadRequestError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, todoList)
}

func (h *Handler) updateList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		responseWithError(c, err)
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responseWithError(c, models.NewBadRequestError(err.Error()))
		return
	}

	var input models.UpdateTodoListInput
	if err := c.BindJSON(&input); err != nil {
		responseWithError(c, models.NewBadRequestError(err.Error()))
		return
	}
	input.Id = listId

	if err := input.Validate(); err != nil {
		responseWithError(c, models.NewBadRequestError(err.Error()))
		return
	}

	if err := h.services.TodoList.Update(userId, input); err != nil {
		responseWithError(c, models.NewInternalServerError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"message": "OK"})
}

func (h *Handler) deleteList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		responseWithError(c, err)
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responseWithError(c, models.NewBadRequestError(err.Error()))
		return
	}

	if err := h.services.TodoList.Delete(userId, id); err != nil {
		responseWithError(c, models.NewInternalServerError(err.Error()))
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"message": "OK"})
}
