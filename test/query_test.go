package test

import (
	"chatross-api/internal/entity"
	"fmt"
	"testing"

	"gorm.io/gorm"
)

func TestSelectExistUser(t *testing.T){
	group := new([]entity.Chat)
	user := []string{"halludba", "michi_lover_test"}
	db.Debug().Where("id IN (?) AND type = 'private'", db.Table("user_chat").Select("chat_id").Where("user_id IN (?)", user)).Take(group)
	fmt.Println(group)
}

func TestGetChatMessage(t *testing.T) {
	chat := new(entity.Chat)
	db.Where("id = ?", 1).Preload("Message", func (db *gorm.DB) *gorm.DB {
		return db.Order("messages.created_at ASC")
	}).Take(chat)
	fmt.Println(chat.Message)
}