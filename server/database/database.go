package database

import (
	"fmt"
	"database/sql"
	"jarvis/constants"
	"crypto/sha256"
)

const (
	host = "roundhouse.proxy.rlwy.net"
	port = 47006
	database = "jarvis"
	username = "postgres"
	password = "KhOTBQkrlwsbIfVExYWnrpRisziefXmn"
)

func Connect() *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, database)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	return db
}

func Close(db *sql.DB) {
	db.Close()
}

func Create_user(db *sql.DB, username string, password string) constants.User {
	id := util.GenerateUUID()
	h := sha256.New()
	h.Write([]byte(password))
	hash := fmt.Sprintf("%x", h.Sum(nil))
	_, err := db.Exec("INSERT INTO users (id, username, password) VALUES ($1, $2, $3)", id, username, hash)
	if err != nil { panic(err) }
	return constants.User{
		Id: id,
		Username: username,
		Apps: []string{},
		Notifications: []string{},
	}
}

func Update_user(db *sql.DB, user constants.User) {
	_, err := db.Exec("UPDATE ONLY users SET apps = $2, notifications = $3, username = $4 WHERE id = $1", user.Id, user.Apps, user.Notifications, user.Username)
	if err != nil { panic(err) }
}

func Update_password(db *sql.DB, id string, old string, new string) {
	h := sha256.New()
	h.Write([]byte(old))
	hash := fmt.Sprintf("%x", h.Sum(nil))
	var resp sql.Row
	if e := db.QueryRow("SELECT hash FROM users WHERE id = $1", old).Scan(&resp); e != nil { panic(e) }
	if resp != hash { panic("Could not authenticate") }
	h.Reset()
	h.Write([]byte(new))
	hash = fmt.Sprintf("%x", h.Sum(nil))
	_, err := db.Exec("UPDATE ONLY users SET hash = $1 WHERE id = $2", hash)
	if err != nil { panic(err) }
}

func Authenticate_user(db *sql.DB, username string, password string) constants.User {
	h := sha256.New()
	h.Write([]byte(password))
	hash := fmt.Sprintf("%x", h.Sum(nil))
	var resp sql.Row
	if e := db.QueryRow("SELECT * FROM users WHERE username = $1 AND password = $2", username, hash).Scan(&resp); e != nil { panic(e) }
	var user constants.User
	err := resp.Scan(&user.Id, &user.Username, &user.Apps, &user.Notifications)
	if err != nil { panic(err) }
	return user
}

func Add_app(db *sql.DB, id string, app string) {
	_, err := db.Exec("UPDATE ONLY users SET apps = apps || '{$2}' WHERE id = $1", id, app)
	if err != nil { panic(err) }
}

func Remove_app(db *sql.DB, id string, app string) {
	_, err := db.Exec("UPDATE ONLY users SET apps = array_remove(apps, $2) WHERE id = $1", id, app)
	if err != nil {
		panic(err)
	}
}

func Get_apps(db *sql.DB, user constants.User) {
	var apps []constants.App
	for _, app := range user.Apps {
		var resp string
		if e := db.QueryRow("SELECT * FROM apps WHERE id = $1", app).Scan(&resp); e != nil { panic(e) }
		if resp != "" {
			apps = append(apps, resp)
		}
	}
}

func Add_notification(db *sql.DB, id string, notification string, title string, content string) {
	_, err := db.Exec("UPDATE ONLY users SET notifications = notifications || '{$2}' WHERE id = $1", id, notification)
	_, err = db.Exec("INSERT INTO notifications (id, title, content) VALUES ($1, $2, $3)", notification, title, content)
	if err != nil {
		panic(err)
	}
}

func Remove_notification(db *sql.DB, id string, notification string) {
	_, err := db.Exec("UPDATE ONLY users SET notifications = array_remove(notifications, $2) WHERE id = $1", id, notification)
	_, err = db.Exec("DELETE FROM notifications WHERE id = $1", notification)
	if err != nil {
		panic(err)
	}
}