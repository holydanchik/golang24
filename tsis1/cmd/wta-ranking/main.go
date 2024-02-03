package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

const PORT = ":8080"

func main() {
	log.Println("starting API server")

	router := mux.NewRouter()

	log.Println("creating routes")
	router.HandleFunc("/players", ListPlayers).Methods("GET")
	router.HandleFunc("/players/{rank}", GetPlayer).Methods("GET")
	router.HandleFunc("/top-players", ListTopPlayers).Methods("GET")
	router.HandleFunc("/health-check", HealthCheck).Methods("GET")

	http.Handle("/", router)

	log.Println("Server started on port", PORT)
	err := http.ListenAndServe(PORT, router)
	log.Fatal(err)
}
