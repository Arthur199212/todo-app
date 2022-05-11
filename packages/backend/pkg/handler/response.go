package handler

import (
	"net/http"
	"todo-app/models"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

func responseWithError(c *gin.Context, err error) {
	logrus.Errorln(err.Error())

	if re, ok := err.(*models.HttpError); ok {
		c.AbortWithStatusJSON(re.Status, map[string]interface{}{
			"message": re.Error(),
		})
		return
	}

	c.AbortWithStatusJSON(http.StatusInternalServerError, map[string]interface{}{
		"message": err.Error(),
	})
}
