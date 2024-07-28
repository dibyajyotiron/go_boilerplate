package repository

import (
	"github.com/go_boilerplate/internal/user/domain"
	"github.com/stretchr/testify/mock"
)

type MockDB struct {
	mock.Mock
}

func (m *MockDB) First(dest interface{}, conds ...interface{}) error {
	args := m.Called(dest, conds[0])
	return args.Error(0)
}

func (m *MockDB) Create(value interface{}) error {
	args := m.Called(value)
	return args.Error(0)
}

type MockUserRepository struct {
	db *MockDB
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{db: &MockDB{}}
}

func (r *MockUserRepository) GetUser(id int) (*domain.User, error) {
	var user domain.User
	err := r.db.First(&user, id)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *MockUserRepository) GetDB() *MockDB {
	return r.db
}

func (r *MockUserRepository) CreateUser(user *domain.User) error {
	return r.db.Create(user)
}
