package usecase

import (
	"chatross-api/internal/entity"
	rerror "chatross-api/internal/helper/error"
	"chatross-api/internal/model"
	"chatross-api/internal/model/wsmodel"
	"chatross-api/internal/repository"
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)


type ChatUsecase struct {
	DB				*gorm.DB
	Log 			*logrus.Logger
	ChatRepository 	*repository.ChatRepository
	Validate 		*validator.Validate
}

func (c *ChatUsecase) GetAllChatFromUser(ctx context.Context, request *model.ChatRequest) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()
}

func (c *ChatUsecase) NewChat(ctx context.Context, request *wsmodel.Message) (int64, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	newUser := &entity.Chat{
		Name: "",
		Type: "private",
		User: []entity.User{
			{ID: request.To},
			{ID: request.From},
		},
	}

	if err := c.ChatRepository.Create(tx, newUser); err != nil {
		c.Log.Error("Failed To Add User")
		return 0, rerror.ErrInternalServer
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.Error("Failed Session Commit")
		return 0, rerror.ErrInternalServer
	}

	return 0, rerror.ErrConflict
	
}