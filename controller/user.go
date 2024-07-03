package controller

import (
	"internnet-backend/handler"
	// "net/http"

	"github.com/go-chi/chi/v5"
)

func CreateUserController(r chi.Router) {
	// r := chi.NewRouter()
	r.Get("/users", handler.GetAllUsers)
	r.Post("/users/login", handler.LoginUser)
	r.Post("/users", handler.CreateUser)
	r.Route("/users/{id}", func(r chi.Router) {
		r.Get("/", handler.GetUserByID)
		r.Put("/", handler.UpdateUser)
		r.Delete("/", handler.DeleteUser)
	})

	//return r
}
