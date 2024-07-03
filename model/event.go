package model

import "time"

type Event struct {
	Id       int      `json:"id"`
	Type     string   `json:"type"`
	Title    string   `json:"title"`
	Location string   `json:"location"`
	Time     time.Time   `json:"time"`
	Userids  []string `json:"userids"`
	Latitude float64 `json"latitude"`
	Longitude float64 `json"longitude"`
}
