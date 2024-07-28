package repository

import "github.com/go_boilerplate/internal/user/domain"

type UserRepository interface {
	GetUser(id int) (*domain.User, error)
	CreateUser(user *domain.User) error
}
