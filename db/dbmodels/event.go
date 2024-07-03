package dbmodels

import models "internnet-backend/model"

import (
	"github.com/lib/pq"
	"time"
)

type Event struct {
	Id       int      `db:"id"`
	Type     string   `db:"type"`
	Title    string   `db:"title"`
	Location string   `db:"location"`
	Time     time.Time   `db:"time"`
	Userids  pq.StringArray `db:"userids"`
	Latitude float64 `db:"latitude"`
	Longitude float64 `db:"longitude"` 
}

func (e Event) ToAPIModel() *models.Event {
	return &models.Event{
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
