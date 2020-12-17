package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/shodhangk/go-auth/models"
)

func Orders(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	v := ctx.Value("user").(*models.Token)
	var user models.User
	_ = models.DB.First(&user, v.UserID)
	var orders []models.Order
	//var users []models.User
	models.DB.Preload("OrderItems.Order").Find(&orders)
	models.DB.Model(&user).Preload("Items").Association("Orders").Find(&orders)
	json.NewEncoder(w).Encode(orders)
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	v := ctx.Value("user").(*models.Token)
	order := &models.Order{}
	json.NewDecoder(r.Body).Decode(order)
	var user models.User
	_ = models.DB.First(&user, v.UserID)
	order.UserID = user.ID
	order.AddItemPrice()
	result := models.DB.Create(&order)
	var errMessage = result.Error

	if result.Error != nil {
		fmt.Println(errMessage)
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusCreated)
	}
	json.NewEncoder(w).Encode(result)
}
