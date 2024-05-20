package router

import (
	"net/http"

	"github.com/alimrndev/go-api/controller"
	"github.com/gorilla/mux"
)

// NewRouter membuat dan mengembalikan router baru
func NewRouter() http.Handler {
	r := mux.NewRouter()

	// Define API version group
	apiV1 := r.PathPrefix("/api").Subrouter()

	// Rute untuk API User
	userController := controller.NewUserController()
	apiV1.HandleFunc("/users", userController.CreateUser).Methods("POST")
	apiV1.HandleFunc("/users", userController.GetAllUsers).Methods("GET")
	apiV1.HandleFunc("/users/{id}", userController.GetUserByID).Methods("GET")
	apiV1.HandleFunc("/users/{id}", userController.UpdateUser).Methods("PUT")
	apiV1.HandleFunc("/users/{id}", userController.DeleteUser).Methods("DELETE")

	// Rute untuk API Menu
	menuController := controller.NewMenuController()
	apiV1.HandleFunc("/menu_items", menuController.CreateMenu).Methods("POST")
	apiV1.HandleFunc("/menu_items", menuController.GetAllMenus).Methods("GET")
	apiV1.HandleFunc("/menu_items/{id}", menuController.GetMenuByID).Methods("GET")
	apiV1.HandleFunc("/menu_items/{id}", menuController.UpdateMenu).Methods("PUT")
	apiV1.HandleFunc("/menu_items/{id}", menuController.DeleteMenu).Methods("DELETE")

	return r
}
