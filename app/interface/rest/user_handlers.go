package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/jedi4z/go-mongodb/app/usecase"
	log "github.com/sirupsen/logrus"
	"net/http"
)

type UserDTO struct {
	ID        string `form:"id" json:"id"`
	FirstName string `form:"first_name" json:"first_name" binding:"required"`
	LastName  string `form:"last_name" json:"last_name" binding:"required"`
	Email     string `form:"email" json:"email" binding:"required"`
}

func toUserDTO(user *usecase.UserUC) *UserDTO {
	return &UserDTO{
		ID:        user.ID,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}

func toUserDTOList(users []*usecase.UserUC) []*UserDTO {
	res := make([]*UserDTO, len(users))

	for i, user := range users {
		res[i] = toUserDTO(user)
	}

	return res
}

type userHandler interface {
	handleNewUser(c *gin.Context)
	handleListUsers(c *gin.Context)
	handleGetUser(c *gin.Context)
	handleUpdateUser(c *gin.Context)
}

func (s service) handleNewUser(c *gin.Context) {
	var userRest UserDTO

	// Binding userRest data
	if err := c.ShouldBindJSON(&userRest); err != nil {
		log.Errorf("error binding user: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := s.userUsecase.RegisterUser(userRest.FirstName, userRest.LastName, userRest.Email)

	// Register a new userRest
	if err != nil {
		log.Errorf("error registering user: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, toUserDTO(user))
}

func (s service) handleListUsers(c *gin.Context) {
	users, err := s.userUsecase.ListUser()
	if err != nil {
		log.Errorf("error getting users: %v", err)
		c.JSON(http.StatusServiceUnavailable, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, toUserDTOList(users))
}

func (s service) handleGetUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "TODO: get a user"})
}

func (s service) handleUpdateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "TODO: update a user"})
}
