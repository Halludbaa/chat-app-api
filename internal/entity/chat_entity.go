package entity

type Chat struct {
	ID 			int64 	`gorm:"column:id;primarykey;autoincrement:true;index"`
	Name 		string 	`gorm:"column:name;default:null"`
	Type		string  `gorm:"column:type;constraint:check:(type in ('private', 'group'))"`
	User 		[]User 	`gorm:"many2many:user_chat"`
	CreatedAt 	int64   `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt 	int64   `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}