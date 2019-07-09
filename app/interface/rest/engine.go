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

type Handler struct {
	userUsecase usecase.UserUsecase
}

func newRequestHandler(ctn *registry.Container) RequestHandler {
	return &Handler{
		userUsecase: ctn.Resolve("user-usecase").(usecase.UserUsecase),
	}
}

func NewEngine(ctn *registry.Container) *gin.Engine {
	r := gin.Default()
	h := newRequestHandler(ctn)

	r.GET("/ping", h.handlePing)
	r.POST("/users", h.handleNewUser)
	r.GET("/users", h.handleListUsers)
	r.GET("/users/{id}", h.handleGetUser)
	r.PUT("/users/{id}", h.handleUpdateUser)

	return r
}
