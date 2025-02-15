package repository

import (
	"chatross-api/internal/entity"

	"gorm.io/gorm"
)

type UserRepository struct {
	Repository[entity.User]
}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) FindByEmail(db *gorm.DB, entity *entity.User, email string) (error) {
	return db.Where("email = ?", email).Take(entity).Error
}

func (r *UserRepository) CountByEmail(db *gorm.DB, email string) (int64, error) {
	var total int64
	err := db.Model(new(entity.User)).Where("email = ?", email).Count(&total).Error
	return total, err
}

func (r *UserRepository) UserWithChat(db *gorm.DB, entity *entity.User) (error) {
	return db.Where("id = ?", entity.ID).Preload("Chat").Take(entity).Error
}
