package models

import "gorm.io/gorm"

type Admin struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	Name     string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
}
