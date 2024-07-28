package service

import (
	"github.com/go_boilerplate/internal/user/domain"
	"github.com/go_boilerplate/internal/user/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUser(id int) (*domain.User, error) {
	return s.repo.GetUser(id)
}

func (s *UserService) CreateUser(user *domain.User) error {
	return s.repo.CreateUser(user)
}
