package repository

import (
	"chatross-api/internal/entity"
	wsmodel "chatross-api/internal/model/wsmodel"

	"gorm.io/gorm"
)


type Chat interface {
	FindFromSenderReceiver(db *gorm.DB, entity *entity.Chat, msg *wsmodel.Message) error
}
type ChatRepository struct {
	Repository[entity.Chat]
}

func NewChatRepository() Chat {
	return &ChatRepository{}
}

func (r *ChatRepository) FindFromSenderReceiver(db *gorm.DB, entity *entity.Chat, msg *wsmodel.Message) error {
	return r.DB.Where("from = ? and to = ?", msg.From, msg.To).Preload("Message").Take(entity).Error
}