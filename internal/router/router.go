package router

import (
	"github.com/go_boilerplate/internal/auth"
	"github.com/go_boilerplate/internal/user/handler"
	"github.com/go_boilerplate/internal/user/repository"
	"github.com/go_boilerplate/internal/user/service"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	userRepo := repository.NewGormUserRepository(db)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r := gin.Default()

	// Public routes
	r.POST("/login", userHandler.Login)

	// Protected routes
	userGroup := r.Group("/user")
	userGroup.Use(auth.AuthenticateUser())
	{
		userGroup.GET("/:id", userHandler.GetUser)
		userGroup.POST("/", userHandler.CreateUser)
	}

	return r
}
