package main

import (
	route "coffeego/route"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	route.Router(router)
	server(router)
}

func server(router *mux.Router) {
	log.Println("Server run on Port 8081")
	log.Fatal(http.ListenAndServe(":8081", handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization", "Cache-Control"}),
		handlers.ExposedHeaders([]string{"Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedOrigins([]string{"*"}))(router),
	))
}
