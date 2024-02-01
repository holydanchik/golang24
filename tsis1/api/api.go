package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Friend struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	// Add more fields as needed
}

var friends = []Friend{
	{Name: "John", Age: 25},
	{Name: "Alice", Age: 28},
	// Add more friends
}

func ListFriends(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(friends)
}

func GetFriend(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	friendName := params["name"]

	for _, friend := range friends {
		if friend.Name == friendName {
			json.NewEncoder(w).Encode(friend)
			return
		}
	}

	http.NotFound(w, r)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("My Web App is healthy! - Author: Your Name"))
}

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/friends", ListFriends).Methods("GET")
	r.HandleFunc("/friends/{name}", GetFriend).Methods("GET")
	r.HandleFunc("/health", HealthCheck).Methods("GET")
}
