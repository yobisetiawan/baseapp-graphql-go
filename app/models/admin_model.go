package models

type Admin struct {
	BaseModel
	Name     string `json:"name" gorm:"type:varchar(255);not null"`
	Email    string `json:"email" gorm:"type:varchar(255);not null"`
	Password string `json:"-" gorm:"type:varchar(255);not null"`
	IsActive bool   `json:"is_active" gorm:"default:false"`
}

func (Admin) TableName() string {
	return "admins"
}
