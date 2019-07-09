package handler

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type Handler struct {
	mongoClient *mongo.Client
}

func (h Handler) HandlePing(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}
