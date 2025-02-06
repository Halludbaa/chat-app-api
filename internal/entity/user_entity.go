package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	// ID 			int64 	`gorm:"primarykey;auto_increment:true;index"`
	ID 			string	`gorm:"column:id;primarykey;index"`
	Email     	string  `gorm:"column:email"`
	CreatedAt 	int64   `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt 	int64   `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	Password  	string 	`gorm:"column:password"`
	Chat 		[]Chat 	`gorm:"many2many:user_chat"`
}

