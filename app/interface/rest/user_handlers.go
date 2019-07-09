package rest

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jedi4z/go-mongodb/app/usecase"
	log "github.com/sirupsen/logrus"
	"net/http"
	"time"
)

type UserRestDTO struct {
	ID        string    `form:"id" json:"id"`
	CreatedAt time.Time `form:"created_at" json:"created_at"`
	FirstName string    `form:"first_name" json:"first_name" binding:"required"`
	LastName  string    `form:"last_name" json:"last_name" binding:"required"`
	Email     string    `form:"email" json:"email" binding:"required"`
}

func toUserRestDTO(user *usecase.UserDTO) *UserRestDTO {
	return &UserRestDTO{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		FirstName: user.FirstName,
		LastName:  user.LastName,
		Email:     user.Email,
	}
}

func toUserRestDTOList(users []*usecase.UserDTO) []*UserRestDTO {
	res := make([]*UserRestDTO, len(users))

	for i, user := range users {
		res[i] = toUserRestDTO(user)
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
	var userRest UserRestDTO

	// Binding userRest data
	if err := c.ShouldBindJSON(&userRest); err != nil {
		log.Errorf("error binding user: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("error with data received: %v", err),
		})
		return
	}

	// Register a new userRest
	user, err := s.userUseCase.RegisterUser(
		userRest.FirstName,
		userRest.LastName,
		userRest.Email,
	)
	if err != nil {
		log.Errorf("impossible to persist a new user: %v", err)
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("impossible to persist a new user: %v", err),
		})
		return
	}

	c.JSON(http.StatusCreated, toUserRestDTO(user))
}

func (s service) handleListUsers(c *gin.Context) {
	users, err := s.userUseCase.ListUser()
	if err != nil {
		log.Errorf("impossible to get users: %v", err)
		c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{
			"error": fmt.Sprintf("impossible to get users: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, toUserRestDTOList(users))
}

func (s service) handleGetUser(c *gin.Context) {
	id := c.Param("id")
	user, err := s.userUseCase.RetrieveAnUser(id)
	if err != nil {
		log.Errorf("error getting the user: %v", err)
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": fmt.Sprintf("user does not exists: %v", err),
		})
		return
	}

	c.JSON(http.StatusOK, toUserRestDTO(user))
}

func (s service) handleUpdateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "TODO: update a user"})
}
