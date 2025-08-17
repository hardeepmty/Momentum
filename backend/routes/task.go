package routes

import (
	"back/controllers"
	"back/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

func TaskRoutes(r *mux.Router){
	r.Handle("/creatework", middleware.AuthMiddleware(http.HandlerFunc(controllers.CreateWorkSpace))).Methods("POST")
	r.Handle("/createtask", middleware.AuthMiddleware(http.HandlerFunc(controllers.CreateTask))).Methods("POST")
	r.Handle("/mytasks", middleware.AuthMiddleware(http.HandlerFunc(controllers.GetUserTasks))).Methods("GET")
}