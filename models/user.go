package models

import (
	"fmt"

	"gorm.io/gorm"
)

//User struct declaration
type User struct {
	gorm.Model
	Name     string
	Email    string `gorm:"type:varchar(100);unique_index"`
	Gender   string `json:"Gender"`
	Password string `json:"Password"`
	Orders   []Order
	Cart     Cart
}

func (u *User) AfterCreate(tx *gorm.DB) (err error) {
	var cart Cart
	fmt.Println("test")
	cart.UserID = u.ID
	result := tx.Create(&cart)
	var errMessage = result.Error

	if errMessage != nil {
		return errMessage
	}
	return
}
