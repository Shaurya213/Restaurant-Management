package models

import "gorm.io/gorm"

type Admin struct {
	gorm.Model

	Name     string `gorm:"type:varchar(100);uniqueIndex;not null"`
	Password string `gorm:"not null"`
}
