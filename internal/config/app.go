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
	messageRepository := repository.NewMessageRepository()
	chatRepository := repository.NewChatRepository()


	// Setup UseCase
	userUseCase := usecase.NewUserUsecase(config.DB, userRepository, config.Validate, config.Log)
	messageUseCase := usecase.NewMessageUsecase(config.DB, messageRepository, config.Validate, config.Log)
	chatUseCase := usecase.NewChatUsecase(config.DB, chatRepository, config.Validate, config.Log)


	// Setup Controller
	authController := controller.NewAuthController(userUseCase)
	websocketController := controller.NewWsController(config.Hub)
	chatController := controller.NewChatController(chatUseCase)
	userController := controller.NewUserController(userUseCase)

	// Setup Middleware
	authMiddleware := middleware.NewAuth(userUseCase)
	config.Hub.UC = &websockets.UseCaseList{
		MessageUsecase: messageUseCase,
		ChatUsecase: chatUseCase,
	}
	
	router := route.RouteConfig{
		App: config.App,
		AuthController: authController,
		AuthMiddleware: authMiddleware,
		ChatController: chatController,
		UserController: userController,
		WebsocketHandler:  websocketController,
		Hub: config.Hub,
	}
	router.Setup()

}