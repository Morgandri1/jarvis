package routes

import (
	"net/http"
	"encoding/json"
	"jarvis/constants"
	"jarvis/database"
)

func Get_apps(w http.ResponseWriter, r *http.Request) {
	var payload constants.SimpleAuthPayload 
    err := json.NewDecoder(r.Body).Decode(&payload)
	user := database.Easy_auth(payload.Username, payload.Password)
	db := database.Connect()
	apps := database.Get_apps(db, user)
}

func Post_apps(w http.ResponseWriter, r *http.Request) {
	
}