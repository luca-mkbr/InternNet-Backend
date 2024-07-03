package dbmodels

import (
	"internnet-backend/model"

	"github.com/lib/pq"
)

type User struct {
	ID        int            `db:"id"`
	FirstName string         `db:"first_name"`
	LastName  string         `db:"last_name"`
	Password  string         `db:"pass_word"`
	Email     string         `db:"email"`
	Company   string         `db:"company"`
	Friends   pq.StringArray `db:"friends"`
	EventIDs  pq.StringArray `db:"event_ids"`
}

func (u User) ToAPIModel() *model.User {
	return &model.User{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
		Password:  u.Password,
		Email:     u.Email,
		Company:   u.Company,
		Friends:   u.Friends,
		EventIDs:  u.EventIDs,
	}
}
