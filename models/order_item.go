package models

type OrderItem struct {
	ID       uint `gorm:"primary_key;auto_increment" json:"id"`
	OrderID  uint `json:"order_id"`
	ItemID   uint `json:"item_id"`
	Quantity uint
	Price    int16
}
