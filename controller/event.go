package controller

import (
	handlers "internnet-backend/handler"

	"github.com/go-chi/chi/v5"
)

func EventRouter(r chi.Router) {
	r.Get("/events/{date}/{type}/{location}", handlers.HandleGetEventsByFilter)
	r.Get("/events", handlers.HandleGetAllEvents)
	r.Post("/events", handlers.HandleCreateEvent)
	r.Route("/events/{id}", func(r chi.Router) {
		r.Get("/", handlers.HandleGetEventById)
		r.Put("/", handlers.HandleUpdateEvent)
		r.Delete("/", handlers.HandleDeleteEvent)
	})

}
