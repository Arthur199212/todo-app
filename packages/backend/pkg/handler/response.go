package handler

import (
	"net/http"
	"todo-app"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

func responseWithError(c *gin.Context, err error) {
	logrus.Errorln(err.Error())

	if re, ok := err.(*todo.RequestError); ok {
		c.AbortWithStatusJSON(re.Status, re.Error())
		return
	}

	c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
}
