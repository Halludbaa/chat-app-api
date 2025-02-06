package test

import (
	"chatross-api/internal/config"
	"chatross-api/internal/delivery/websockets"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)
var (
	app *gin.Engine
	db 	*gorm.DB
	validate *validator.Validate
	hub *websockets.Hub

)
func init(){
	config.LoadEnv()
	gin.SetMode(gin.ReleaseMode)
	app = gin.New()
	db = config.NewDatabase()
	validate = validator.New()
	hub = websockets.NewHub()

	config.Boostrap(&config.BoostrapConfig{
		App: app,
		Validate: validate,
		DB: db,
		Hub: hub,
	})
}
