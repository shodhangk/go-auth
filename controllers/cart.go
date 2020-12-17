package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/shodhangk/go-auth/models"
	"gorm.io/gorm/clause"
)

type addItemsParams struct {
	Items []models.CartItem
}

func CartItems(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	v := ctx.Value("user").(*models.Token)
	var cart models.Cart
	_ = models.DB.First(&cart, "user_id = ?", v.UserID)
	json.NewEncoder(w).Encode(cart.ItemList())
}

func CurrentUserID(r *http.Request) (userID uint) {
	ctx := r.Context()
	v := ctx.Value("user").(*models.Token)
	return v.UserID
}

func CurrentUser(r *http.Request) (user models.User) {
	userID := CurrentUserID(r)
	models.DB.First(&user, userID)
	return
}

func UserCart(r *http.Request) (c models.Cart) {
	user := CurrentUser(r)
	models.DB.Model(&user).Preload("Cart")
	c = user.Cart
	return
}

func AddToCart(w http.ResponseWriter, r *http.Request) {
	var items addItemsParams
	json.NewDecoder(r.Body).Decode(&items)
	var cart models.Cart
	userID := CurrentUserID(r)
	_ = models.DB.First(&cart, "user_id = ?", userID)
	for i := 0; i < len(items.Items); i++ {
		items.Items[i].CartID = cart.ID
		fmt.Println(cart.ID)
	}
	if err := models.DB.Create(&items.Items).Error; err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}
	w.WriteHeader(http.StatusCreated)
	models.DB.Preload(clause.Associations).First(&cart, "user_id = ?", userID)
	json.NewEncoder(w).Encode(cart.ItemList())
}

func DeleteItem(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	cartID := params["item_id"]
	// Convert `orderId` string param to uint64
	id64, _ := strconv.ParseUint(cartID, 10, 64)
	// Convert uint64 to uint
	idToDelete := uint(id64)
	user.
		models.DB.Where("item_id = ? AND cart_id =?", idToDelete, UserCart(r)).Delete(&models.CartItem{})
	w.WriteHeader(http.StatusNoContent)
}

func UpdateCart(w http.ResponseWriter, r *http.Request) {
	var cartItem models.CartItem
	json.NewDecoder(r.Body).Decode(&cartItem)
	models.DB.Model(&models.CartItem{}).Where("item_id = ? AND cart_id =?", cartItem.ItemID, CurrentUserID(r)).Update("quantiy", cartItem.Quantity)
}
