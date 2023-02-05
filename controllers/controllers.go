package controllers

import (
	"api-go-rest/database"
	"api-go-rest/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func ShowAllProducts(w http.ResponseWriter, r *http.Request) {
	var p []models.Product
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}
