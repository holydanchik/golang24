package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	handlers "github.com/holydanchik/golang24/tsis1/pkg/wta-ranking/handlers"
)

const PORT = ":8080"

func main() {
	log.Println("starting API server")

	router := mux.NewRouter()

	log.Println("creating routes")
	router.HandleFunc("/players", handlers.ListPlayers).Methods("GET")
	router.HandleFunc("/players/{rank}", handlers.GetPlayer).Methods("GET")
	router.HandleFunc("/top-players", handlers.ListTopPlayers).Methods("GET")
	router.HandleFunc("/health-check", handlers.HealthCheck).Methods("GET")

	http.Handle("/", router)

	log.Println("Server started on port", PORT)
	err := http.ListenAndServe(PORT, router)
	log.Fatal(err)
}
