package model

type User struct {
	ID        int      `json:"ID"`
	FirstName string   `json:"firstName"`
	LastName  string   `json:"lastName"`
	Password  string   `json:"password"`
	Email     string   `json:"email"`
	Company   string   `json:"company"`
	Friends   []string `json:"friends"`
	EventIDs  []string `json:"eventIDs"`
}
