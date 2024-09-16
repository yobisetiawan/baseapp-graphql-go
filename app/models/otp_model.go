package models

import "time"

type Otp struct {
	BaseModel
	Email     string     `json:"email" gorm:"type:varchar(255);"`
	OtpKey    string     `json:"otp_key" gorm:"type:varchar(20);"`
	ExpiredAt *time.Time `json:"expired_at"`
	Purpose   string     `json:"purpose" gorm:"type:varchar(50); default:'general';index"`
	Status    string     `json:"status" gorm:"type:varchar(30); default:'active';index"`
}

func (Otp) TableName() string {
	return "otps"
}
