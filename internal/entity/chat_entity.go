package entity

type Chat struct {
	ID 				int64 		`json:"id" gorm:"column:id;primarykey;autoincrement:true;index"`
	Name 			string 		`json:"name,omitempty" gorm:"column:name;default:null"`
	Type			string  	`json:"type" gorm:"column:type;constraint:check:(type in ('private', 'group'))"`

	User1ID			string 		`json:"-,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User1			User		`json:"user_1,omitempty"`

	User2ID			string		`json:"-,omitempty" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	User2			User		`json:"user_2,omitempty"`
	
	User 			[]User 		`json:"user,omitempty" gorm:"many2many:user_chat"`
	Message			[]Message 	`json:"message" gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	CreatedAt 		int64   	`json:"created_at" gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt 		int64   	`json:"updated_at" gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}