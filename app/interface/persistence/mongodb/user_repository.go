package mongodb

import (
	"context"
	"github.com/jedi4z/go-mongodb/app/domain/model"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

const (
	dbName         = "go_mongodb"
	collectionUser = "users"
)

var (
	mongoConnectionUri = os.Getenv("MONGO_CONNECTION_URI")
)

type userRepository struct {
	mongoClient *mongo.Client
	users       map[string]*User
}

func NewUserRepository() *userRepository {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoConnectionUri))
	if err != nil {
		panic(err)
	}

	return &userRepository{
		mongoClient: client,
		users:       map[string]*User{},
	}
}

func (r *userRepository) FindAll() ([]*model.User, error) {
	// Here's an array in which you can store the decoded documents
	results := make([]*model.User, len(r.users))

	ctx := context.TODO()
	collection := r.mongoClient.Database(dbName).Collection(collectionUser)

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Errorf("failed to get users: %v", err)
		return nil, err
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(ctx) {
		// create a value into which the single document can be decoded
		var elem User

		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		user := model.NewUser(elem.ID, elem.FirstName, elem.LastName, elem.Email)

		results = append(results, user)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	if err := cur.Close(ctx); err != nil {
		log.Fatal(err)
	}

	return results, nil
}

func (r *userRepository) FindByEmail(email string) (*model.User, error) {
	var elem User

	ctx := context.TODO()
	filter := bson.M{"email": email}
	collection := r.mongoClient.Database(dbName).Collection(collectionUser)

	if err := collection.FindOne(ctx, filter).Decode(&elem); err != nil {
		return nil, err
	}

	user := model.NewUser(elem.ID, elem.FirstName, elem.LastName, elem.Email)

	return user, nil
}

func (r *userRepository) Save(user *model.User) error {
	collection := r.mongoClient.Database(dbName).Collection(collectionUser)
	_, err := collection.InsertOne(context.TODO(), &User{
		ID:        user.GetID(),
		FirstName: user.GetFirstName(),
		LastName:  user.GetLastName(),
		Email:     user.GetEmail(),
	})

	return err
}

type User struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
}
