package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type pingHandler interface {
	handlePing(c *gin.Context)
}

func (h Handler) handlePing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
