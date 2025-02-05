package converter

import (
	"chatross-api/internal/entity"
	"chatross-api/internal/model"
)

func UserToResponse(user *entity.User) *model.UserResponse{
	return &model.UserResponse{
		Username: user.ID,
		Email: user.Email,
		CreateAt: user.CreatedAt,
		UpdateAt: user.UpdatedAt,
	}
} 