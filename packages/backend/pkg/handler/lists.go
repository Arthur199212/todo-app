package handler

import (
	"net/http"
	"strconv"
	"todo-app"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

func (h *Handler) getAllLists(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		responseWithError(c, todo.NewRequestError(http.StatusUnauthorized, err))
		return
	}

	// todo: pagination

	todoLists, err := h.services.TodoList.GetAll(userId)
	if err != nil {
		responseWithError(c, todo.NewRequestError(http.StatusInternalServerError, err))
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
		responseWithError(c, todo.NewRequestError(http.StatusUnauthorized, err))
		return
	}

	var input todoListInput
	if err := c.BindJSON(&input); err != nil {
		responseWithError(c, todo.NewRequestError(http.StatusBadRequest, err))
		return
	}

	if err := input.Validate(); err != nil {
		responseWithError(c, todo.NewRequestError(http.StatusBadRequest, err))
		return
	}

	listId, err := h.services.TodoList.Create(userId, input.Title)
	if err != nil {
		responseWithError(c, todo.NewRequestError(http.StatusInternalServerError, err))
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"id": listId})
}

func (h *Handler) getListById(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		responseWithError(c, todo.NewRequestError(http.StatusUnauthorized, err))
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responseWithError(c, todo.NewRequestError(http.StatusBadRequest, err))
		return
	}

	todoList, err := h.services.TodoList.GetById(userId, id)
	if err != nil {
		responseWithError(c, todo.NewRequestError(http.StatusBadRequest, err))
		return
	}

	c.JSON(http.StatusOK, todoList)
}

func (h *Handler) updateList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		responseWithError(c, todo.NewRequestError(http.StatusUnauthorized, err))
		return
	}

	listId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responseWithError(c, todo.NewRequestError(http.StatusBadRequest, err))
		return
	}

	var input todo.UpdateTodoListInput
	if err := c.BindJSON(&input); err != nil {
		responseWithError(c, todo.NewRequestError(http.StatusBadRequest, err))
		return
	}
	input.Id = listId

	if err := input.Validate(); err != nil {
		responseWithError(c, todo.NewRequestError(http.StatusBadRequest, err))
		return
	}

	if err := h.services.TodoList.Update(userId, input); err != nil {
		responseWithError(c, todo.NewRequestError(http.StatusInternalServerError, err))
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"message": "OK"})
}

func (h *Handler) deleteList(c *gin.Context) {
	userId, err := getUserId(c)
	if err != nil {
		responseWithError(c, todo.NewRequestError(http.StatusUnauthorized, err))
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		responseWithError(c, todo.NewRequestError(http.StatusBadRequest, err))
		return
	}

	if err := h.services.TodoList.Delete(userId, id); err != nil {
		responseWithError(c, todo.NewRequestError(http.StatusInternalServerError, err))
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{"message": "OK"})
}
