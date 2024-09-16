package models

import "time"

type User struct {
	BaseModel
	Name             string     `json:"name" gorm:"type:varchar(255);not null"`
	Email            string     `json:"email" gorm:"type:varchar(255);not null;uniqueIndex"`
	Password         string     `json:"-" gorm:"type:varchar(255);not null"`
	EmailVerifiedAt  *time.Time `json:"email_verified_at" gorm:"default: null"`
	Status           string     `json:"status" gorm:"type:varchar(255);default:'active';index"`
	Role             string     `json:"role" gorm:"type:varchar(255); default:'customer';index"`
	AvatarURL        string     `json:"avatar_url" gorm:"type:varchar(255); null;"`
	MarkForDeletedAt *time.Time `json:"mark_for_deleted_at" gorm:"default: null"`
	NameOld          string     `json:"-" gorm:"type:varchar(255);null"`
	EmailOld         string     `json:"-" gorm:"type:varchar(255);null"`
}

func (User) TableName() string {
	return "users"
}
