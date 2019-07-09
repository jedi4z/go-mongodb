package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/jedi4z/go-mongodb/app/registry"
	"github.com/jedi4z/go-mongodb/app/usecase"
)

type RequestHandler interface {
	userHandler
	pingHandler
}

type service struct {
	userUsecase usecase.UserUsecase
}

func newService(ctn *registry.Container) RequestHandler {
	return &service{
		userUsecase: ctn.Resolve("user-usecase").(usecase.UserUsecase),
	}
}

func NewRestEngine(ctn *registry.Container) *gin.Engine {
	r := gin.Default()
	s := newService(ctn)

	r.GET("/ping", s.handlePing)
	r.POST("/users", s.handleNewUser)
	r.GET("/users", s.handleListUsers)
	r.GET("/users/{id}", s.handleGetUser)
	r.PUT("/users/{id}", s.handleUpdateUser)

	return r
}
