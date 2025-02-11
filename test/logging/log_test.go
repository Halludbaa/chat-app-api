package test

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
)

func TestLog(t *testing.T){
	logger := logrus.New()
	// logger.SetLevel(logrus.)
	logger.SetFormatter(&logrus.JSONFormatter{})


	file, _ := os.OpenFile("app.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)

	logger.Trace("Hello World!")
	logger.Debug("Hello World!")
	logger.Info("Hello World!")
	logger.Warn("Hello World!")
	logger.Error("Hello World!")
}