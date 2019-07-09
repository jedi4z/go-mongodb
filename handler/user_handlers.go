package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jedi4z/go-mongodb/common"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func (h Handler) HandleNewUser(c *gin.Context) {
	userCollection := h.mongoClient.Database(common.DBName).Collection(common.CollectionUser)
	res, err := userCollection.InsertOne(c, bson.M{"name": "pi", "value": 3.14159})
	if err != nil {
		log.Errorf("error creating a user: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{"error": "Error!"})
	}

	c.JSON(http.StatusCreated, res)
}

func (h Handler) HandleListUsers(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "TODO: list users"})
}

func (h Handler) HandleGetUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "TODO: get a user"})
}

func (h Handler) HandleUpdateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "TODO: update a user"})
}
