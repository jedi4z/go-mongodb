package rest

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
	if err := h.userUsecase.RegisterUser("jesusdiazbc@gmail.com"); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "OK"})
}

func (h Handler) handleListUsers(c *gin.Context) {
	users, err := h.userUsecase.ListUser()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
	}

	c.JSON(http.StatusOK, users)
}

func (h Handler) handleGetUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "TODO: get a user"})
}

func (h Handler) handleUpdateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "TODO: update a user"})
}
