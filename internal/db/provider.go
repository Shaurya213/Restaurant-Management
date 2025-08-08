package db

import "gorm.io/gorm"

// ProvideDB exposes the initialized *gorm.DB for DI.
// Call InitMySQL() once (in main) before using this.
func ProvideDB() *gorm.DB {
	return DB
}
