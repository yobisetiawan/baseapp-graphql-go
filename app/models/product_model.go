package models

import "time"

type Product struct {
	BaseModel
	Name        string    `json:"name" gorm:"type:varchar(255);not null"`
	Description string    `json:"description"`
	StartDate   time.Time `json:"start_date"`
	EndDate     time.Time `json:"end_date"`
	SellerId    string    `json:"seller_id" gorm:"not null"`
	Seller      *User     `json:"seller,omitempty" gorm:"foreignKey:SellerId;references:ID"`
	Price       float64   `json:"price" gorm:"type:decimal(10,2);not null;default:0"`
	Type        string    `json:"type" gorm:"type:varchar(50);default:'regular';index"`
	Status      string    `json:"status" gorm:"type:varchar(50);default:'draft';index"`
}

func (Product) TableName() string {
	return "products"
}
