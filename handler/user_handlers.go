package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/jedi4z/go-mongodb/common"
	"github.com/jedi4z/go-mongodb/model"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"net/http"
)

func (h Handler) HandleNewUser(c *gin.Context) {
	var user model.User

	// Binding user data
	if err := c.ShouldBindJSON(&user); err != nil {
		log.Errorf("error binding user: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	collection := h.mongoClient.Database(common.DBName).Collection(common.CollectionUser)

	// Insert the new user
	_, err := collection.InsertOne(c, user)
	if err != nil {
		log.Errorf("error creating a user: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{"error": "Error!"})
		return
	}

	c.JSON(http.StatusCreated, user)
}

func (h Handler) HandleListUsers(c *gin.Context) {
	// Here's an array in which you can store the decoded documents
	var results []*model.User

	collection := h.mongoClient.Database(common.DBName).Collection(common.CollectionUser)

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := collection.Find(c, bson.D{})
	if err != nil {
		log.Errorf("error retrieving users: %s", err.Error())
		c.AbortWithStatusJSON(http.StatusServiceUnavailable, gin.H{"error": "Error!"})
		return
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(c) {

		// create a value into which the single document can be decoded
		var elem model.User
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	if err := cur.Close(c); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, results)
}

func (h Handler) HandleGetUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "TODO: get a user"})
}

func (h Handler) HandleUpdateUser(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "TODO: update a user"})
}
