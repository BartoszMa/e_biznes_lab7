package scopes

import (
	"gorm.io/gorm"
)

func CartWithItemsAndProducts() func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Preload("CartItems.Product")
	}
}
