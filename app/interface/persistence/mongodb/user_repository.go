package mongodb

import (
	"context"
	"github.com/jedi4z/go-mongodb/app/domain/model"
	"go.mongodb.org/mongo-driver/mongo"
)

const (
	DBName         = "go_mongodb"
	CollectionUser = "users"
)

type userRepository struct {
	mongoClient *mongo.Client
	users       map[string]*User
}

func (r *userRepository) FindAll() ([]*model.User, error) {
	return nil, nil
}

func (r *userRepository) FindByEmail(email string) (*model.User, error) {
	return nil, nil
}

func (r *userRepository) Save(user *model.User) error {
	collection := r.mongoClient.Database(DBName).Collection(CollectionUser)
	_, err := collection.InsertOne(context.TODO(), user)

	return err
}

type User struct {
	ID    string
	Email string
}
