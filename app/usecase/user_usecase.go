package usecase

import (
	"github.com/google/uuid"
	"github.com/jedi4z/go-mongodb/app/domain/model"
	"github.com/jedi4z/go-mongodb/app/domain/repository"
	"github.com/jedi4z/go-mongodb/app/domain/service"
)

type UserUsecase interface {
	ListUser() ([]*User, error)
	RegisterUser(firstName, lastName, email string) error
}

type userUsecase struct {
	repo    repository.UserRepository
	service *service.UserService
}

func NewUserUsecase(repo repository.UserRepository, service *service.UserService) *userUsecase {
	return &userUsecase{
		repo:    repo,
		service: service,
	}
}

func (u *userUsecase) ListUser() ([]*User, error) {
	users, err := u.repo.FindAll()

	if err != nil {
		return nil, err
	}

	return toUser(users), nil
}

func (u *userUsecase) RegisterUser(firstName, lastName, email string) error {
	uid, err := uuid.NewRandom()

	if err != nil {
		return err
	}

	if err := u.service.Duplicated(email); err != nil {
		return err
	}

	user := model.NewUser(uid.String(), firstName, lastName, email)
	if err := u.repo.Save(user); err != nil {
		return err
	}

	return nil
}

type User struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
}

func toUser(users []*model.User) []*User {
	res := make([]*User, len(users))

	for i, user := range users {
		res[i] = &User{
			ID:        user.GetID(),
			FirstName: user.GetFirstName(),
			LastName:  user.GetLastName(),
			Email:     user.GetEmail(),
		}
	}

	return res
}
