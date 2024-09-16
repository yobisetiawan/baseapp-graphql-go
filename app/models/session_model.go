package models

import "time"

type Session struct {
	BaseModel
	AdminId      string    `json:"admin_id" gorm:"default:null"`
	Admin        Admin     `gorm:"foreignKey:AdminId;references:ID"`
	UserId       string    `json:"user_id" gorm:"default:null"`
	User         User      `gorm:"foreignKey:UserId;references:ID"`
	IsActive     bool      `json:"is_active" gorm:"default:false"`
	ExpiredAt    time.Time `json:"expired_at" `
	RefreshToken string    `json:"refresh_token" gorm:"type:varchar(500);default:null;index"`
}

func (Session) TableName() string {
	return "sessions"
}
