package router

import (
	"net/http"

	"github.com/alimrndev/go-api/controller"
	"github.com/gorilla/mux"
)

// NewRouter membuat dan mengembalikan router baru
func NewRouter() http.Handler {
	r := mux.NewRouter()
	userController := controller.NewUserController()

	// Rute untuk API Create User
	r.HandleFunc("/api/users", userController.CreateUser).Methods("POST")

	// Rute untuk API Read User
	r.HandleFunc("/api/users", userController.GetAllUsers).Methods("GET")

	// Rute untuk API Read User by ID
	r.HandleFunc("/api/users/{id}", userController.GetUserByID).Methods("GET")

	// Rute untuk API Update User
	r.HandleFunc("/api/users/{id}", userController.UpdateUser).Methods("PUT")

	// Rute untuk API Delete User
	r.HandleFunc("/api/users/{id}", userController.DeleteUser).Methods("DELETE")

	return r
}
