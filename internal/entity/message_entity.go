package entity

type Message struct {
	ID 			int64 	`gorm:"column:id;primarykey;autoincrement:true;index"`
	ChatID		int64 	`gorm:"column:chat_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	From 		string  `gorm:"column:from"`
	To			string 	`gorm:"column:to"`
	CreatedAt 	int64   `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt 	int64   `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}