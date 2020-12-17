package models

import (
	"github.com/shodhangk/go-auth/models/jsons"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Cart struct {
	gorm.Model
	UserID    uint `json:"user_id"`
	CartItems []CartItem
	Items     []Item `gorm:"many2many:cart_items"`
}

func (cart *Cart) ItemList() (cartJSON jsons.CartJson) {
	DB.Preload(clause.Associations).First(&cart)
	cartJSON.Items = []jsons.CartItemJson{}
	for i := 0; i < len(cart.CartItems); i++ {
		cartJSON.Items = append(cartJSON.Items, jsons.CartItemJson{
			cart.Items[i].ID,
			cart.Items[i].Name,
			cart.CartItems[i].Quantity,
			cart.Items[i].Price})
	}
	return
}
