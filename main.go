package main

import (
	"github.com/gorilla/mux"
	"go_todolist/handlers"
	"go_todolist/models"
	"log"
	"net/http"
)

func main() {
	// Initialize the database
	models.InitDB()

	// Initialize the router
	r := mux.NewRouter()

	// Apply logging middleware to all routes
	r.Use(handlers.LoggingMiddleware)

	// Public routes
	r.HandleFunc("/login", handlers.Login).Methods("POST")
	r.HandleFunc("/register", handlers.Register).Methods("POST")

	// Protected routes
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/todolist", handlers.GetToDoList).Methods("GET")
	api.HandleFunc("/todolist", handlers.CreateToDo).Methods("POST")
	api.Use(handlers.JWTAuthMiddleware)

	// Start the server
	log.Println("Server started at :8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
