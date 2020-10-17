package coffeego

import (
	app "coffeego/app"
	"net/http"

	"github.com/gorilla/mux"
)

// Router Should Exported
func Router(router *mux.Router) {
	router.Handle("/api/products", http.HandlerFunc(app.ProductListRequest)).Methods("GET")
	router.Handle("/api/product/{id}", http.HandlerFunc(app.ProductRead)).Methods("GET")

	router.Handle("/api/users", http.HandlerFunc(app.UserListRequest)).Methods("GET")
	router.Handle("/api/me", http.HandlerFunc(app.CurrentUserRequest)).Methods("GET")

}
