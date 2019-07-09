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
	userUseCase usecase.UserUseCase
}

func newService(ctn *registry.Container) RequestHandler {
	return &service{
		userUseCase: ctn.Resolve("user-use-case").(usecase.UserUseCase),
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
