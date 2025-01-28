package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID 			int64 	`gorm:"primarykey;auto_increment:true;index"`
	Name 		string	`gorm:"column:name"`
	Email     	string  `gorm:"column:email"`
	CreatedAt 	int64   `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt 	int64   `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
	Password  	string 	`gorm:"column:password"`
}