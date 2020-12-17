package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/shodhangk/go-auth/models"
)

func Items(w http.ResponseWriter, r *http.Request) {
	var items []models.Item
	models.DB.Find(&items)
	fmt.Println(items)
	json.NewEncoder(w).Encode(items)
}

func Item(w http.ResponseWriter, r *http.Request) {
	var item models.Item
	models.DB.First(&item)
	json.NewEncoder(w).Encode(item)
}
