package handler

import (
	"encoding/json"
	"internnet-backend/data"
	"internnet-backend/db"
	"internnet-backend/db/dbmodels"
	"internnet-backend/model"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

func apiToUserDbModel(u *model.User) *dbmodels.User {
	return &dbmodels.User{
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

func GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := data.GetAllUsers()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	d, err := data.GetUserByID(idInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(d)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	dbUser := apiToUserDbModel(&user)
	createdUser, err := data.CreateUser(*dbUser)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(createdUser)
}

// Changed for Swift compatibility
func LoginUser(w http.ResponseWriter, r *http.Request) {
	var LoginInput model.LoginInput

	err := json.NewDecoder(r.Body).Decode(&LoginInput)

	if err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
	}

	if LoginInput.Email == "" || LoginInput.Password == "" {
		http.Error(w, "Email and password are required", http.StatusBadRequest)
		return
	}

	user := model.User{}
	database := db.GetDbConnection()
	err = database.QueryRow("SELECT id , pass_word FROM users WHERE email = $1", LoginInput.Email).Scan(&user.ID, &user.Password)
	if err != nil {
		//http.Error(w, "User not found", http.StatusUnauthorized)
		http.Error(w, "0000", http.StatusUnauthorized)
		return
	}

	if user.Password != LoginInput.Password {
		http.Error(w, "Invalid Password", http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	// json.NewEncoder(w).Encode(model.LoginResponse{
	// 	Message: "Login successful",
	// 	UserID:  user.ID,
	// })
	json.NewEncoder(w).Encode(user.ID)

}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	user := model.User{}
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	id := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
	user.ID = idInt
	dbUser := apiToUserDbModel(&user)
	updatedUser, err := data.UpdateUser(*dbUser)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedUser)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	user, err := data.DeleteUser(idInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
