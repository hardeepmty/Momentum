package main

import (
	"back/config"
	"back/routes"
	"log"
	"net/http"
)

func main() {
	config.ConnectDB()

	router := routes.RegisterRoutes()

	log.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000",router))
}