package model


type UserResponse struct {
	ID 		int64	`json:"id,omitempty"`
	Name	string	`json:"name,omitempty"`
	Email 	string 	`json:"email,omitempty"`
	CreateAt int64 `json:"create_at,omitempty"`
	UpdateAt int64 `json:"update_at,omitempty"`
}

type RegisterUserRequest struct {
	Name string `json:"name" validate:"required,max=100" `
	Password string `json:"password" validate:"required,max=255"`
	Email string `json:"email" validate:"required,max=100" `
}

type LoginUserRequest struct {
	Email string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
