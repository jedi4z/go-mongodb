package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/jedi4z/go-mongodb/app/usecase"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type UserRest struct {
	ID        string `form:"id" json:"id"`
	FirstName string `form:"first_name" json:"first_name" binding:"required"`
	LastName  string `form:"last_name" json:"last_name" binding:"required"`
	Email     string `form:"email" json:"email" binding:"required"`
}

func toUserRest(user *usecase.UserUC) *UserRest {
	return &UserRest{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}

func toUserRestList(users []*usecase.UserUC) []*UserRest {
	res := make([]*UserRest, len(users))

	for i, user := range users {
		res[i] = toUserRest(user)
	}

	return res
}

type userHandler interface {
	handleNewUser(c *gin.Context)
	handleListUsers(c *gin.Context)
	handleGetUser(c *gin.Context)
	handleUpdateUser(c *gin.Context)
}

func (h Handler) handleNewUser(c *gin.Context) {
	var userRest UserRest

	// Binding userRest data
	if err := c.ShouldBindJSON(&userRest); err != nil {
		log.Errorf("error binding user: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userUsecase.RegisterUser(userRest.FirstName, userRest.LastName, userRest.Email)

	// Register a new userRest
	if err != nil {
		log.Errorf("error registering user: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, toUserRest(user))
}

func (h Handler) handleListUsers(c *gin.Context) {
	users, err := h.userUsecase.ListUser()
	if err != nil {
		log.Errorf("error getting users: %v", err)
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, toUserRestList(users))
}

func (h Handler) handleGetUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "TODO: get a user"})
}

func (h Handler) handleUpdateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "TODO: update a user"})
}
