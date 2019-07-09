package rest

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type userHandler interface {
	handleNewUser(c *gin.Context)
	handleListUsers(c *gin.Context)
	handleGetUser(c *gin.Context)
	handleUpdateUser(c *gin.Context)
}

type User struct {
	ID        string `form:"id" json:"id"`
	FirstName string `form:"first_name" json:"first_name" binding:"required"`
	LastName  string `form:"last_name" json:"last_name" binding:"required"`
	Email     string `form:"email" json:"email" binding:"required"`
}

func (h Handler) handleNewUser(c *gin.Context) {
	var user User

	// Binding user data
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Errorf("error binding user: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Register a new user
	if err := h.userUsecase.RegisterUser(user.FirstName, user.LastName, user.Email); err != nil {
		log.Errorf("error registering user: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, user)
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
