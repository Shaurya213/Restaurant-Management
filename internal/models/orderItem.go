package models

import "gorm.io/gorm"

type OrderItem struct {
	gorm.Model
	OrderID    uint
	MenuItemID uint
	Quantity   int     `gorm:"not null"`
	Price      float64 `gorm:"not null"`
}
