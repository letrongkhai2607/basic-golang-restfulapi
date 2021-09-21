package model

import "github.com/jinzhu/gorm"

type Transfer struct {
	gorm.Model
	From_User_ID int`gorm:"not null" json:"from_user_id" validate:"required"`
	To_User_ID int`gorm:"not null" json:"to_user_id" validate:"required"`
	Amount int`gorm:"not null" json:"amount" validate:"required"`
}
