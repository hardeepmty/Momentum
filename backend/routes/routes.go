package routes

import (
	"github.com/gorilla/mux"
)

func RegisterRoutes() *mux.Router {
	r := mux.NewRouter()

	authRouter := r.PathPrefix("/auth").Subrouter()
	AuthRoutes(authRouter)

	apiRouter := r.PathPrefix("/api").Subrouter()
	TaskRoutes(apiRouter) 

	return r
}