package test

import (
	"chatross-api/internal/config"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)
var (
	app *gin.Engine
	db 	*gorm.DB
	// validate *validator.Validate
	// hub *websockets.Hub

)
func init(){
	config.LoadEnv()
	// gin.SetMode(gin.ReleaseMode)
	// app = gin.New()
	db = config.NewDatabase()
	// validate = validator.New()
	// hub = websockets.NewHub()

	// config.Boostrap(&config.BoostrapConfig{
	// 	App: app,
	// 	Validate: validate,
	// 	DB: db,
	// 	Hub: hub,
	// })
}
