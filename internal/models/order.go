package models

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Customer string
	Phone    string
	Items    []OrderItem `gorm:"foreignKey:OrderID"`
	Total    float64     `gorm:"not null"`
	IsPaid   bool        `gorm:"default:false"`
	IsServed bool        `gorm:"default:false"`
}
