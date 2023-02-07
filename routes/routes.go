package routes

import (
	"api-go-rest/controllers"
	"api-go-rest/middleware"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(middleware.ContentTypeMiddleware)
	r.HandleFunc("/", controllers.Home)
	r.HandleFunc("/list", controllers.ShowAllProducts)
	r.HandleFunc("/product/{id}", controllers.FindById).Methods("Get")
	r.HandleFunc("/product", controllers.Save).Methods("Post")
	log.Fatal(http.ListenAndServe(":8000", r))
}
