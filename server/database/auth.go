package database

import (
	"./database"
)

func Easy_auth(username string, password string) bool {
	db := database.Connect()
	user := database.Authenticate_user(db, username, password)
	db.close()
	return user
}