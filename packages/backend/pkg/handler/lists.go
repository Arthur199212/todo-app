package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) getAllLists(c *gin.Context) {
	// just for test
	userId, ok := c.Get(userCtx)
	if !ok {
		newErrorResponse(c, http.StatusUnauthorized, "something went wrong")
		return
	}
	c.JSON(http.StatusOK, map[string]interface{}{
		"message": "test went well",
		"userId":  userId,
	})
}

func (h *Handler) createList(c *gin.Context) {}

func (h *Handler) getListById(c *gin.Context) {}

func (h *Handler) updateList(c *gin.Context) {}

func (h *Handler) deleteList(c *gin.Context) {}
