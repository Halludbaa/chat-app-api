package entity

type User struct {
	ID 			string	`json:"id,omitempty" gorm:"column:id;primarykey;index"`
	Email     	string  `json:"email,omitempty" gorm:"column:email"`
	CreatedAt 	int64   `json:"created_at,omitempty" gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt 	int64   `json:"updated_at,omitempty" gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	Password  	string 	`json:"-" gorm:"column:password"`
	Chat 		[]Chat 	`json:"chat,omitempty" gorm:"many2many:user_chat"`
}

