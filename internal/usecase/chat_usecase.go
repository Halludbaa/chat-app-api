package usecase

import (
	"chatross-api/internal/model"
	"chatross-api/internal/model/wsmodel"
	"chatross-api/internal/repository"
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)


type ChatUsecase struct {
	DB	*gorm.DB
	Log *logrus.Logger
	ChatRepository *repository.ChatRepository
	Validate *validator.Validate
}

func (c *ChatUsecase) GetAllChatFromUser(ctx context.Context, request *model.ChatRequest) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()
}

func (c *ChatUsecase) NewChat(request *wsmodel.Message) {

	
}