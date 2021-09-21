package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique_index;not null" json:"username" validate:"required,min=8,max=32"`
	Password string `gorm:"not null" json:"password" validate:"required,min=8,max=32"`

	Email    string `gorm:"unique_index" json:"email" validate:"email,min=8,max=32"`
	Names    string `json:"names"`
	Phone    string `json:"phone"`
	Balance int `json:"balance"`
}