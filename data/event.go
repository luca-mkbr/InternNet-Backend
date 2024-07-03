package data

import (
	// "errors"
	"fmt"

	"internnet-backend/db"
	"internnet-backend/db/dbmodels"
)

// Gets all events in the database
func GetAllEvents() ([]dbmodels.Event, error) {

	database := db.GetDbConnection()
	var events []dbmodels.Event
	err := database.Select(&events, "SELECT * FROM events")
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	return events, nil
}

// Gets a specific event based on event id
func GetEventById(id int) (dbmodels.Event, error) {
	database := db.GetDbConnection()
	var event dbmodels.Event
	err := database.Get(&event, "SELECT * FROM events WHERE id=$1", id)
	if err != nil {
		return dbmodels.Event{}, err
	}
	return event, nil
}

// Gets events using a filter
func GetEventsByFilter( /*eventDate time.Time, */ eventType string, eventLocation string) ([]dbmodels.Event, error) {

	// fmt.Println(eventDate)

	database := db.GetDbConnection()
	var events []dbmodels.Event

	if eventType == "any" && eventLocation == "any" {
		err := database.Select(&events, "SELECT * FROM events")
		if err != nil {
			return nil, err
		}
	} else if eventType == "any" && eventLocation != "any" {
		fmt.Println("GOT HERE 1")
		err := database.Select(&events, "SELECT * FROM events WHERE location=$1", eventLocation)
		if err != nil {
			return nil, err
		}
	} else if eventType != "any" && eventLocation == "any" {
		fmt.Println("GOT HERE 2")
		fmt.Println(eventType)
		fmt.Println(eventLocation)
		err := database.Select(&events, "SELECT * FROM events WHERE type=$1", eventType)
		if err != nil {
			return nil, err
		}
	} else {
		err := database.Select(&events, "SELECT * FROM events WHERE type=$1 AND location=$2", eventType, eventLocation)
		if err != nil {
			return nil, err
		}
	}

	// if eventDate == "any" {
	// 	return events, nil
	// }

	var returnEvents []dbmodels.Event
	// for _, event := range events {
	// 	if event.Time == eventDate {
	// 		returnEvents = append(returnEvents, event)
	// 	}
	// }

	return returnEvents, nil
}

// Creates an event
func CreateEvent(event dbmodels.Event) (dbmodels.Event, error) {

	fmt.Println(event.Time)

	database := db.GetDbConnection()

	rows, err := database.NamedQuery(`INSERT INTO events ("type", title, location, time, userids, latitude, longitude) 
	VALUES (:type, :title, :location, :time, :userids, :latitude, :longitude) RETURNING id`, event)
	if err != nil {
		return dbmodels.Event{}, err
	}
	// fmt.Println(rows)
	var id int
	if rows.Next() {
		rows.Scan(&id)
		event.Id = id
		rows.Close()
	}
	return event, nil
}

// Updates an event based on event id
func UpdateEvent(event dbmodels.Event) (dbmodels.Event, error) {
	database := db.GetDbConnection()
	rst, err := database.NamedExec(`UPDATE events SET "type"=:type, title=:title, location=:location,
	time=:time, userids=:userids, latitude:=latitude, longitude:=longitude WHERE id=:id`, event)
	if err != nil {
		return dbmodels.Event{}, err
	}
	n, _ := rst.RowsAffected()
	if n == 0 {
		return dbmodels.Event{}, fmt.Errorf("no items found with id:%d", event.Id)
	}
	return event, nil
}

// Deletes an event based on event id
func DeleteEvent(id int) (int, error) {
	database := db.GetDbConnection()
	rst, err := database.Exec("DELETE FROM events WHERE id=$1", id)
	if err != nil {
		return -1, err
	}
	n, _ := rst.RowsAffected()
	if n == 0 {
		return -1, fmt.Errorf("no items found with id:%d", id)
	}

	return id, nil
}
