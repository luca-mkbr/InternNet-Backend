package main

import (
	"fmt"
	"internnet-backend/controller"
	"net/http"
	
	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	controller.CreateUserController(r)
	controller.EventRouter(r)
	// controller.UserRouter(r)
	fmt.Println("Server listening on port 8080.")
	http.ListenAndServe("localhost:8080", r)
}