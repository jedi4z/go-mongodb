package usecase

import (
	"github.com/google/uuid"
	"github.com/jedi4z/go-mongodb/app/domain/model"
	"github.com/jedi4z/go-mongodb/app/domain/repository"
	"github.com/jedi4z/go-mongodb/app/domain/service"
)

type UserUC struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
}

func toUserUC(user *model.User) *UserUC {
	return &UserUC{
		ID:        user.GetID(),
		FirstName: user.GetFirstName(),
		LastName:  user.GetLastName(),
		Email:     user.GetEmail(),
	}
}

func toUserUCList(users []*model.User) []*UserUC {
	res := make([]*UserUC, len(users))

	for i, user := range users {
		res[i] = toUserUC(user)
	}

	return res
}

type UserUsecase interface {
	ListUser() ([]*UserUC, error)
	RegisterUser(firstName, lastName, email string) (*UserUC, error)
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

func (u *userUsecase) ListUser() ([]*UserUC, error) {
	users, err := u.repo.FindAll()

	if err != nil {
		return nil, err
	}

	return toUserUCList(users), nil
}

func (u *userUsecase) RegisterUser(firstName, lastName, email string) (*UserUC, error) {
	uid, err := uuid.NewRandom()

	if err != nil {
		return nil, err
	}

	if err := u.service.Duplicated(email); err != nil {
		return nil, err
	}

	user := model.NewUser(uid.String(), firstName, lastName, email)
	if err := u.repo.Save(user); err != nil {
		return nil, err
	}

	return toUserUC(user), nil
}
