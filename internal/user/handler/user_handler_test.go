package handler

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go_boilerplate/internal/user/domain"
	"github.com/go_boilerplate/internal/user/repository"
	"github.com/go_boilerplate/internal/user/service"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetUser(t *testing.T) {
	mockRepo := repository.NewMockUserRepository()
	userService := service.NewUserService(mockRepo)
	userHandler := NewUserHandler(userService)

	user := &domain.User{ID: 1, Name: "John Doe", Email: "john@example.com"}

	mockRepo.GetDB().On("First", &domain.User{}, user.ID).Run(func(args mock.Arguments) {
		arg := args.Get(0).(*domain.User)
		*arg = *user
	}).Return(nil)

	r := gin.Default()
	r.GET("/user/:id", userHandler.GetUser)

	req, _ := http.NewRequest("GET", "/user/1", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockRepo.GetDB().AssertExpectations(t)
}

func TestCreateUser(t *testing.T) {
	mockRepo := repository.NewMockUserRepository()
	userService := service.NewUserService(mockRepo)
	userHandler := NewUserHandler(userService)

	user := &domain.User{ID: 1, Name: "John Doe", Email: "john@example.com"}

	mockRepo.GetDB().On("Create", user).Return(nil)

	r := gin.Default()
	r.POST("/user", userHandler.CreateUser)

	jsonStr := `{"id": 1, "name": "John Doe", "email": "john@example.com"}`
	req, _ := http.NewRequest("POST", "/user", strings.NewReader(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	mockRepo.GetDB().AssertExpectations(t)
}
