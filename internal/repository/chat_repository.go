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

func NewChatRepository() *ChatRepository {
	return &ChatRepository{}
}

func (r *ChatRepository) FindFromSenderReceiver(db *gorm.DB, entity *entity.Chat, user []string) error {
	return 	db.Debug().Where(
		"(user1_id = ? AND user2_id = ?) OR (user1_id = ? AND user2_id = ?)", user[0], user[1],
		user[1], user[0]).First(entity).Error
}

func (r *ChatRepository) GetChatWithMessage(db *gorm.DB, entity *entity.Chat, chatID int64) error {
	return db.Where("id = ?", chatID).Preload("Message").
									Preload("User").
									Preload("User1").
									Preload("User2").Take(entity).Error
}
