package models

type CartItem struct {
	ID       uint `gorm:"primary_key;auto_increment" json:"id"`
	CartID   uint `json:"cart_id"`
	ItemID   uint `json:"item_id"`
	Quantity uint
}
