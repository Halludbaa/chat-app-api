package repository

import (
	"chatross-api/internal/entity"

	"gorm.io/gorm"
)

type Message interface {
	FindMessageFromChatID(db *gorm.DB, entity *entity.Message, userID string) error
}

type MessageRepository struct {
	Repository[entity.Message]

}

func NewMessageRepository() *MessageRepository {
	return &MessageRepository{}
}

func (r *MessageRepository) FindMessageFromChatID(db *gorm.DB, entity *entity.Message, chatID string) error {
	return r.DB.Where("chat_id = ?", chatID).Take(entity).Error
}

