package constants

type User struct {
	Id string `json:"id"`
	Username string `json:"username"`
	Apps []string `json:"apps"`
	Notifications []string `json:"notifications"`
}

type AuthResponse struct {
	Authenticated bool `json:"authenticated"`
	Token string `json:"token"`
	User User `json:"user"`
}