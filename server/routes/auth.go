package routes

import (
	"net/http"
	"encoding/json"
	"jarvis/constants"
	"jarvis/database"
)

func Auth(w http.ResponseWriter, r *http.Request) {
	var user_payload constants.AuthRequest 
    err := json.NewDecoder(r.Body).Decode(&user_payload)
    if err != nil {
        // return HTTP 400 bad request
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
    db := database.Connect()
    var user constants.User
	switch user_payload.Method {
		case "LOGIN":
			// check if user exists
			user = database.Authenticate_user(db, user_payload.Username, user_payload.Password)
		case "REGISTER":
			// create user
			user = database.Create_user(db, user_payload.Username, user_payload.Password)
	}
	res := constants.Response{
		Status: 0,
		Data: constants.AuthResponse{
			Authenticated: true,
			Token: "",
			User: user,
		},
		Error: constants.JarvisError{
			Error: "",
			Message: "",
		},
	}
	json.NewEncoder(w).Encode(res)
	db.Close()
}