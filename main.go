package main

import (
	"log"
	"net/http"

	"github.com/drew138/graphics-api/routes"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterRoutes(r)
	log.Println("Server started, running on port 8080.")
	if err := http.ListenAndServe("127.0.0.1:8080", r); err != nil {
		log.Fatal("Server failed to start: ", err.Error())
	}
}
