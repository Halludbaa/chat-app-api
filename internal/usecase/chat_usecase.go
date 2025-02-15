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

type Chat interface {
	NewChat(ctx context.Context, request *wsmodel.Message) (int64, error)
	GetAllChatFromUser(ctx context.Context, request *model.ChatRequest)
}

type ChatUsecase struct {
	db				*gorm.DB
	log 			*logrus.Logger
	chatRepository 	*repository.ChatRepository
	validate 		*validator.Validate
}

func NewChatUsecase(db *gorm.DB, chatRepository *repository.ChatRepository, validate *validator.Validate, log *logrus.Logger) *ChatUsecase {
	return &ChatUsecase{
		db: db,
		log: log,
		chatRepository: chatRepository,
		validate: validate,
	}
}

func (c *ChatUsecase) GetChatwithMessage(ctx context.Context, chatID int64) (*entity.Chat, error) {
	tx := c.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	chat := new(entity.Chat)
	if err := c.chatRepository.GetChatWithMessage(tx, chat, chatID); err != nil {
		c.log.WithField("error", err).Error("Failure in Database!")
		return nil, rerror.ErrNotFound
	}
	
	// make simple handler if the requester is not the chat member

	
	if err := tx.Commit().Error; err != nil {
		c.log.WithError(err).Error("Failed Session Commit")
		return nil, rerror.ErrInternalServer
	}

	return chat, nil
}

func (c *ChatUsecase) GetChatID(ctx context.Context, request *wsmodel.Message) (int64, error) {
	tx := c.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	user := []string{request.From, request.To}
	chat := new(entity.Chat)

	if _ = c.chatRepository.FindFromSenderReceiver(tx, chat, user); chat.ID != 0{
		return chat.ID, nil
	}

	chatID, err := c.newChat(ctx, request)
	if err != nil || chatID == 0 {
		c.log.WithField("error", err).Error("Failure in Database!")
		return 0, err
	}

	if err := tx.Commit().Error; err != nil {
		c.log.WithError(err).Error("Failed Session Commit")
		return 0, err
	}

	return chatID, nil

	 
}



func (c *ChatUsecase) newChat(ctx context.Context, request *wsmodel.Message) (int64, error) {
	tx := c.db.WithContext(ctx).Begin()
	defer tx.Rollback()

	newChat := &entity.Chat{
		Name: "",
		Type: "private",
		User1ID: request.From,
		User2ID: request.To,
		User: []entity.User{{ ID: request.From }, {ID: request.To}},
	}

	if err := c.chatRepository.Create(tx, newChat); err != nil {
		c.log.WithError(err).Error("Failed To Add User")
		return 0, rerror.ErrInternalServer
	}

	if err := tx.Commit().Error; err != nil {
		c.log.WithError(err).Error("Failed Session Commit")
		return 0, rerror.ErrInternalServer
	}

	return newChat.ID, nil
	
}