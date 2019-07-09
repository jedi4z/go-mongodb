package service

import (
	"fmt"
	"github.com/jedi4z/go-mongodb/app/domain/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (s *UserService) Duplicated(email string) error {
	user, _ := s.repo.FindByEmail(email)

	if user != nil {
		return fmt.Errorf("%s already exists", email)
	}

	return nil
}
