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

type JarvisApplication struct {
	Id string `json:"id"`
	Name string `json:"name"`
	Owner string `json:"owner"`
	Route string `json:"route"`
	Frequency int `json:"frequency"`
	Priority int `json:"priority"`
}