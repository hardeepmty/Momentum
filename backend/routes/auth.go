package routes

import (
	"back/controllers"

	"github.com/gorilla/mux"
)

func AuthRoutes(r *mux.Router) {
	r.HandleFunc("/signup", controllers.Signup).Methods("POST")

	r.HandleFunc("/login", controllers.Login).Methods("POST")
}
