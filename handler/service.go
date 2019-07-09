package handler

import "go.mongodb.org/mongo-driver/mongo"

func NewService(mongoClient *mongo.Client) RequestHandler {
	return &Handler{mongoClient: mongoClient}
}
