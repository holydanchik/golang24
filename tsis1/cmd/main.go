package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/holydanchik/golang24/tsis1/api"
)

const port = ":8080"

func main() {
	log.Println("starting API server")
	//create a new router
	router := mux.NewRouter()
	log.Println("creating routes")

	api.SetupRoutes(router)

	http.Handle("/", router)

	log.Println("Server started on port", port)
	//start and listen to requests
	http.ListenAndServe(port, router)

}
