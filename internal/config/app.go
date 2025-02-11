package config

import (
	"chatross-api/internal/delivery/http/controller"
	"chatross-api/internal/delivery/http/middleware"
	"chatross-api/internal/delivery/http/route"
	"chatross-api/internal/delivery/websockets"
	"chatross-api/internal/repository"
	"chatross-api/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type BoostrapConfig struct {
	DB 			*gorm.DB
	App			*gin.Engine
	Log			*logrus.Logger
	Validate 	*validator.Validate
	Hub 		*websockets.Hub
}
func Boostrap(config *BoostrapConfig) {
	// Setup Repository
	userRepository := repository.NewUserRepository()

	// Setup UseCase
	userUseCase := usecase.NewUserUsecase(config.DB, userRepository, config.Validate, config.Log)

	// Setup Controller
	authController := controller.NewAuthController(userUseCase)
	websocketController := controller.NewWsController(config.Hub)


	// Setup Middleware
	authMiddleware := middleware.NewAuth(userUseCase)

	
	router := route.RouteConfig{
		App: config.App,
		AuthController: authController,
		AuthMiddleware: authMiddleware,
		WebsocketHandler:  websocketController,
		Hub: config.Hub,
	}
	router.Setup()

}