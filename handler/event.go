package handler

import (
	"encoding/json"
	"fmt"
	"internnet-backend/data"
	"internnet-backend/db/dbmodels"
	models "internnet-backend/model"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func HandleGetAllEvents(w http.ResponseWriter, r *http.Request) {
	data, err := data.GetAllEvents()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var apiData []models.Event
	for i, _ := range data{
		apiData = append(apiData, *data[i].ToAPIModel())
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(apiData)
}

func HandleGetEventById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	d, err := data.GetEventById(idInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(d.ToAPIModel())
}

func HandleGetEventsByFilter(w http.ResponseWriter, r *http.Request) {
	// eventDate := chi.URLParam(r, "date")
	eventType := chi.URLParam(r, "type")
	eventLocation := chi.URLParam(r, "location")

	fmt.Println(eventType)
	fmt.Println(eventLocation)

	data, err := data.GetEventsByFilter(/*eventDate, */eventType, eventLocation)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(data)
}

func apiToDbModel(e *models.Event) *dbmodels.Event {
	return &dbmodels.Event{
		Id:       e.Id,
		Type:     e.Type,
		Title:    e.Title,
		Location: e.Location,
		Time:     e.Time,
		Userids:  e.Userids,
		Latitude: e.Latitude,
		Longitude: e.Longitude,
	}
}

func HandleCreateEvent(w http.ResponseWriter, r *http.Request) {
	event := models.Event{}
	err := json.NewDecoder(r.Body).Decode(&event)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	dbEvent := apiToDbModel(&event)
	createPub, err := data.CreateEvent(*dbEvent)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createPub.ToAPIModel())
}

func HandleUpdateEvent(w http.ResponseWriter, r *http.Request) {
	event := models.Event{}
	err := json.NewDecoder(r.Body).Decode(&event)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	id := chi.URLParam(r, "id")
	eventId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	event.Id = eventId
	dbEvent := apiToDbModel(&event)
	updatedEvent, err := data.UpdateEvent(*dbEvent)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedEvent.ToAPIModel())
}

func HandleDeleteEvent(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	deletedEventId, err := data.DeleteEvent(idInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(deletedEventId)
}
