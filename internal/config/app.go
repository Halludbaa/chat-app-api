package config

import (
	"chatross-api/internal/delivery/http/controller"
	"chatross-api/internal/delivery/http/middleware"
	"chatross-api/internal/delivery/http/route"
	"chatross-api/internal/repository"
	"chatross-api/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type BoostrapConfig struct {
	DB 			*gorm.DB
	App			*gin.Engine
	Validate 	*validator.Validate
}
func Boostrap(config *BoostrapConfig) {
	// Setup Repository
	userRepository := repository.NewUserRepository()

	// Setup UseCase
	userUseCase := usecase.NewUserUseCase(config.DB, userRepository, config.Validate)

	// Setup Controller
	authController := controller.NewAuthController(userUseCase)

	// Setup Middleware
	authMiddleware := middleware.NewAuth(userUseCase)

	
	router := route.RouteConfig{
		App: config.App,
		AuthController: authController,
		AuthMiddleware: authMiddleware,
	}
	router.Setup()

}