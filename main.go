package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/jedi4z/go-mongodb/handler"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

func initEngine(rh handler.RequestHandler) {
	r := gin.Default()

	// Ping endpoint
	r.GET("/ping", rh.HandlePing)

	// User endpoints
	users := r.Group("/users")
	{
		users.POST("", rh.HandleNewUser)
		users.GET("", rh.HandleListUsers)
		users.GET("/{id}", rh.HandleGetUser)
		users.PUT("/{id}", rh.HandleUpdateUser)
	}

	if err := r.Run(); err != nil {
		panic(err)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	mongoConnectionUri := os.Getenv("MONGO_CONNECTION_URI")

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoConnectionUri))
	if err != nil {
		panic(err)
	}

	requestHandler := handler.NewService(client)

	initEngine(requestHandler)
}
