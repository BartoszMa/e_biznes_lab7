package models

import "gorm.io/gorm"

type Payment struct {
	gorm.Model
	IsPayed bool    `json:"is_payed"`
	Price   float32 `json:"price"`
}
