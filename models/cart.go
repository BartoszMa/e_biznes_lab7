package models

import "gorm.io/gorm"

type Cart struct {
	gorm.Model
	CartItems []CartItem `json:"cart_items" gorm:"foreignKey:CartID"`
}

type CartItem struct {
	gorm.Model
	CartID    uint    `json:"cart_id"`
	ProductID uint    `json:"product_id"`
	Product   Product `gorm:"foreignKey:ProductID" json:"product"`
}
