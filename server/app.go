package main

import (
	"fmt"
 	"net/http"
  	"github.com/gorilla/mux"
	"jarvis/routes"
)

const Addr = ":80"

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/auth", routes.Auth).Methods("POST")
	srv := &http.Server{
		Addr: Addr,
	}
	fmt.Println("Server starting on port", Addr)
	srv.ListenAndServe()
}