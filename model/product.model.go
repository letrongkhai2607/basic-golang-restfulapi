package model

import "github.com/jinzhu/gorm"

type Product struct {
	gorm.Model
	Title       string `gorm:"not null" json:"title" validate:"required,min=3,max=32"`
	Description string `gorm:"not null" json:"description" validate:"required,min=3,max=32"`
	Price 		int `gorm:""not null" json:"price" validate:"required""`
	Amount      int    `gorm:"not null" json:"amount" validate:"required,min=3"`
}