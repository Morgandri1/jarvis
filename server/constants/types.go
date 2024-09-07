package constants

type JarvisError struct {
	Error string `json:"error"`
	Message string `json:"message"`
}

type Response struct {
	Status int `json:"status"`
	Data any `json:"data"`
	Error JarvisError `json:"error"`
}