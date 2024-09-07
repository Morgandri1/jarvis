package constants

type AuthRequest struct {
	Method string `json: "method"`// LOGIN, REGISTER
	Username string `json: "username"`
	Password string `json: "password"`
}