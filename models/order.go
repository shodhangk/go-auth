package models

import "gorm.io/gorm"

type Order struct {
	gorm.Model
	UserID     uint   `json:"user_id"`
	Items      []Item `gorm:"many2many:order_items;"`
	TotalPrice int16  `json:"amount"`
}

func (or *Order) AddItemPrice() {
	var item Item
	for i := 0; i < len(or.Items); i++ {
		DB.First(&item, or.Items[i].ID)
		or.Items[i].Price = item.Price
	}
}
