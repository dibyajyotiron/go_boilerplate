package service

import (
	"testing"

	"github.com/go_boilerplate/internal/user/domain"
	"github.com/go_boilerplate/internal/user/repository"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestUserService(t *testing.T) {
	mockRepo := repository.NewMockUserRepository()
	service := NewUserService(mockRepo)

	user := &domain.User{ID: 1, Name: "John Doe", Email: "john@example.com"}

	// Mock the CreateUser method
	mockRepo.GetDB().On("Create", user).Return(nil)
	err := service.CreateUser(user)
	assert.NoError(t, err)
	mockRepo.GetDB().AssertExpectations(t)

	// Mock the GetUser method
	mockRepo.GetDB().On("First", &domain.User{}, user.ID).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*domain.User)
		*arg = *user
	}).Return(nil)
	fetchedUser, err := service.GetUser(user.ID)
	assert.NoError(t, err)
	assert.Equal(t, user, fetchedUser)
	mockRepo.GetDB().AssertExpectations(t)
}
