package usecase

import (
	"chatross-api/internal/entity"
	rerror "chatross-api/internal/helper/error"
	"chatross-api/internal/model/converter"
	"chatross-api/internal/model/wsmodel"
	"chatross-api/internal/repository"
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type MessageUsecase struct {
	DB	*gorm.DB
	Log *logrus.Logger
	MessageRepository *repository.MessageRepository
	Validate *validator.Validate

}

func NewMessageUsecase(db *gorm.DB, msgRepository *repository.MessageRepository, validate *validator.Validate, log *logrus.Logger) *MessageUsecase {
	return &MessageUsecase{
		DB: db,
		Log: log,
		MessageRepository: msgRepository,
		Validate: validate,
	}
}

func (u *MessageUsecase) Store(ctx context.Context, request *wsmodel.Message) (*wsmodel.Message, error) {
	tx := u.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	newUser := &entity.Message{
		ChatID: request.ChatID,
		From: request.From,
		To: request.To,
		Content: request.Content,
	}

	if err := u.MessageRepository.Create(tx, newUser); err != nil {
		u.Log.Error("Failed To Add User")
		return nil, rerror.ErrInternalServer
	}

	if err := tx.Commit().Error; err != nil {
		u.Log.Error("Failed Session Commit")
		return nil, rerror.ErrInternalServer
	}

	return converter.MsgToWsMsg(newUser), nil
}