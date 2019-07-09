package repository

import "github.com/jedi4z/go-mongodb/app/domain/model"

type UserRepository interface {
	FindAll() ([]*model.User, error)
	FindOne(id string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	Save(*model.User) error
}
