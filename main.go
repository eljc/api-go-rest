package main

import (
	"api-go-rest/database"
	"api-go-rest/routes"
)

func main() {
	database.ConectDB()
	routes.HandleRequest()
}
