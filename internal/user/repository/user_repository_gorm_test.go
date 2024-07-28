package repository

import (
	"testing"

	"github.com/go_boilerplate/internal/user/domain"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGormUserRepository(t *testing.T) {
	mockDB := new(MockDB)
	repo := &MockUserRepository{db: mockDB}

	user := &domain.User{ID: 1, Name: "John Doe", Email: "john@example.com"}

	// Mock the Create method
	mockDB.On("Create", user).Return(nil)
	err := repo.CreateUser(user)
	assert.NoError(t, err)
	mockDB.AssertExpectations(t)

	// Mock the First method
	mockDB.On("First", &domain.User{}, 1).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*domain.User)
		*arg = *user
	}).Return(nil)
	fetchedUser, err := repo.GetUser(1)
	assert.NoError(t, err)
	assert.Equal(t, user, fetchedUser)
	mockDB.AssertExpectations(t)
}
