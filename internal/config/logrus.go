package config

import (
	"os"

	"github.com/sirupsen/logrus"
)

func NewLogger() *logrus.Logger{
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{})
	// logger.SetLevel(logrus.InfoLevel)
	file, _ := os.OpenFile(os.Getenv("LOG_FILE"), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	logger.SetOutput(file)

	return logger
}