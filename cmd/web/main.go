package main

import (
	"chatross-api/internal/config"
	"chatross-api/internal/delivery/websockets"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)


func init(){
	config.LoadEnv()
}


func main(){
	db := config.NewDatabase()
	validate := validator.New()
	app := gin.New()
	hub := websockets.NewHub()

	go hub.Run()

	config.Boostrap(&config.BoostrapConfig{
		App: app,
		Validate: validate,
		DB: db,
		Hub: hub,
	})

	err := app.Run(fmt.Sprintf(":%s", os.Getenv("WEB_PORT")))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}



}
