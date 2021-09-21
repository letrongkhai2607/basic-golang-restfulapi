package model

import "github.com/jinzhu/gorm"

type Recharge struct {
	gorm.Model
	To_User_ID int`gorm:"not null" json:"to_user_id" validate:"required"`
	Amount int`gorm:"not null" json:"amount" validate:"required"`
}