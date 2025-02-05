package model


type UserResponse struct {
	// ID 		int64	`json:"id,omitempty"`
	Username 	string	`json:"username,omitempty"`
	Email 		string 	`json:"email,omitempty"`
	CreateAt 	int64 	`json:"create_at,omitempty"`
	UpdateAt 	int64 	`json:"update_at,omitempty"`
}

type RegisterUserRequest struct {
	Username 	string `json:"username" validate:"required,max=100" `
	Password 	string `json:"password" validate:"required,max=255"`
	Email 		string `json:"email,omitempty"`
}

type LoginUserRequest struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
