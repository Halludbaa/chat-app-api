package main

import (
	"chatross-api/internal/config"
	"chatross-api/internal/delivery/websockets"
	"fmt"
	"os"

	"github.com/go-playground/validator/v10"
)


func init(){
	config.LoadEnv()
}


func main(){
	app := config.NewGin()
	db := config.NewDatabase()
	validate := validator.New()
	log := config.NewLogger()
	hub := websockets.NewHub(log, nil)
	go hub.Run()
	defer log.Fatal("App Was Stopped!")

	log.Info("App Is Running!")
	config.Boostrap(&config.BoostrapConfig{
		App: app,
		Validate: validate,
		DB: db,
		Log: log,
		Hub: hub,
	})
	
	err := app.Run(fmt.Sprintf(":%s", os.Getenv("WEB_PORT")))
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	


}
