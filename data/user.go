package data

import (
	"fmt"
	"internnet-backend/db"
	"internnet-backend/db/dbmodels"

	"github.com/google/uuid"
)

// var users []model.User = []model.User{
// 	{
// 		ID:        1,
// 		FirstName: "x",
// 		LastName:  "x",
// 		Password:  "x",
// 		Email:     "x",
// 		Company:   "x",
// 		Friends:   []int{},
// 		// 	{
// 		// 	ID: 			2,
// 		// 	FirstName:		"x",
// 		// 	LastName:		"x",
// 		// 	Password:		"x",
// 		// 	Email:			"x",
// 		// 	Company:		"x",
// 		// 	Friends:		[]model.User{},
// 		// 	EventIDs:		[]string,
// 		// }
// 		EventIDs: []string{},
// 	},
// }

func GetAllUsers() ([]dbmodels.User, error) {
	database := db.GetDbConnection()
	var users []dbmodels.User
	err := database.Select(&users, "SELECT* FROM users")
	if err != nil {
		return nil, err
	}
	return users, nil
}

// func GetAllUsers() ([]model.User, error) {
// 	return users, nil
// 	// return "implementation later"
// }

func CreateUser(user dbmodels.User) (dbmodels.User, error) {
	database := db.GetDbConnection()
	user.ID = int(uuid.New().ID() / 1000)
	_, err := database.NamedExec(`INSERT INTO users (first_name, last_name, pass_word, email, company, friends, event_ids) 
	VALUES (:first_name, :last_name, :pass_word, :email, :company, :friends, :event_ids)`, user)
	if err != nil {
		return dbmodels.User{}, err
	}
	return user, nil
}

// func CreateUser(user model.User) (model.User, error) {
// 	user.ID = int(uuid.New().ID())
// 	users = append(users, user)
// 	return user, nil
// }

func GetUserByID(id int) (dbmodels.User, error) {
	database := db.GetDbConnection()
	var user dbmodels.User
	err := database.Get(&user, "SELECT* FROM users WHERE id=$1", id)
	if err != nil {
		return dbmodels.User{}, err
	}
	return user, nil
}

// func GetUserByID(id int) (model.User, error) {
// 	for _, user := range users {
// 		if user.ID == id {
// 			return user, nil
// 		}
// 	}
// 	return model.User{}, errors.New("User not found")
// }

func UpdateUser(user dbmodels.User) (dbmodels.User, error) {
	database := db.GetDbConnection()
	rst, err := database.NamedExec(`UPDATE users SET first_name=:first_name, last_name=:last_name,
	 pass_word=:pass_word, email=:email, company=:company, friends=:friends, event_ids=:event_ids WHERE id=:id`, user)

	if err != nil {
		return dbmodels.User{}, err
	}
	n, _ := rst.RowsAffected()
	if n == 0 {
		return dbmodels.User{}, fmt.Errorf("no items found with id:%d", user.ID)
	}
	return user, nil
}

// func UpdateUser(user model.User) (model.User, error) {
// 	for i, usr := range users {
// 		if usr.ID == user.ID {
// 			users[i] = user
// 			return user, nil
// 		}
// 	}
// 	return model.User{}, errors.New("User not found")
// }

func DeleteUser(id int) (int, error) {
	database := db.GetDbConnection()
	rst, err := database.Exec("DELETE FROM users WHERE id=$1", id)
	if err != nil {
		return -1, err
	}
	n, _ := rst.RowsAffected()
	if n == 0 {
		return -1, fmt.Errorf("no items found with id:%d", id)
	}

	return id, nil
}

// func DeleteUser(id int) (int, error) {
// 	for i, usr := range users {
// 		if usr.ID == id {
// 			users = append(users[:i], users[i+1:]...)
// 			return id, nil
// 		}
// 	}
// 	return 0, errors.New("User not found")
// }
