package usecase

import (
	"github.com/google/uuid"
	"github.com/jedi4z/go-mongodb/app/domain/model"
	"github.com/jedi4z/go-mongodb/app/domain/repository"
	"github.com/jedi4z/go-mongodb/app/domain/service"
	"time"
)

type UserDTO struct {
	ID        string
	CreatedAt time.Time
	FirstName string
	LastName  string
	Email     string
}

func toUserDTO(user *model.User) *UserDTO {
	return &UserDTO{
		ID:        user.GetID(),
		CreatedAt: user.GetCreatedAt(),
		FirstName: user.GetFirstName(),
		LastName:  user.GetLastName(),
		Email:     user.GetEmail(),
	}
}

func toUserDTOList(users []*model.User) []*UserDTO {
	res := make([]*UserDTO, len(users))

	for i, user := range users {
		res[i] = toUserDTO(user)
	}

	return res
}

type UserUseCase interface {
	ListUser() ([]*UserDTO, error)
	RetrieveAnUser(id string) (*UserDTO, error)
	RegisterUser(firstName, lastName, email string) (*UserDTO, error)
}

type userUseCase struct {
	repository repository.UserRepository
	service    *service.UserService
}

func NewUserUseCase(repo repository.UserRepository, service *service.UserService) UserUseCase {
	return &userUseCase{
		repository: repo,
		service:    service,
	}
}

func (u *userUseCase) ListUser() ([]*UserDTO, error) {
	users, err := u.repository.FindAll()
	if err != nil {
		return nil, err
	}

	return toUserDTOList(users), nil
}

func (u *userUseCase) RetrieveAnUser(id string) (*UserDTO, error) {
	user, err := u.repository.FindOne(id)
	if err != nil {
		return nil, err
	}

	return toUserDTO(user), nil
}

func (u *userUseCase) RegisterUser(firstName, lastName, email string) (*UserDTO, error) {
	uid, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	if err := u.service.Duplicated(email); err != nil {
		return nil, err
	}

	user := model.NewUser(
		uid.String(),
		time.Now().UTC(),
		firstName,
		lastName,
		email,
	)

	if err := u.repository.Save(user); err != nil {
		return nil, err
	}

	return toUserDTO(user), nil
}
