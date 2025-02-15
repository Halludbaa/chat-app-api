package entity

type Message struct {
	ID 			int64 	`json:"id" gorm:"column:id;primarykey;autoincrement:true;index"`
	ChatID		int64 	`json:"chat_id" gorm:"column:chat_id;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	From 		string  `json:"from" gorm:"column:from"`
	To			string 	`json:"to" gorm:"column:to"`
	Content 	string 	`json:"content" gorm:"column:content"`
	CreatedAt 	int64   `json:"created_at" gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt 	int64   `json:"updated_at" gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}