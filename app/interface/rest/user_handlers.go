package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type userHandler interface {
	handleNewUser(c *gin.Context)
	handleListUsers(c *gin.Context)
	handleGetUser(c *gin.Context)
	handleUpdateUser(c *gin.Context)
}

func (h Handler) handleNewUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "TODO: get a user"})
}

func (h Handler) handleListUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "TODO: get a user"})
}

func (h Handler) handleGetUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "TODO: get a user"})
}

func (h Handler) handleUpdateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "TODO: update a user"})
}
