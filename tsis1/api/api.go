// api/api.go
package api

import (
	"encoding/json"
	"net/http"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
)

// Player represents a tennis player's information
type Player struct {
	Rank              int    `json:"rank"`
	Name              string `json:"name"`
	Region            string `json:"region"`
	Age               int    `json:"age"`
	TournamentsPlayed int    `json:"tournaments_played"`
	Points            int    `json:"points"`
}

var players = []Player{
	{Rank: 4, Name: "Jessica Pegula", Region: "USA", Age: 29, TournamentsPlayed: 21, Points: 5705},
	{Rank: 10, Name: "Karolina Muchova", Region: "CZE", Age: 27, TournamentsPlayed: 14, Points: 3520},
	{Rank: 5, Name: "Elena Rybakina", Region: "KAZ", Age: 24, TournamentsPlayed: 18, Points: 5688},
	{Rank: 1, Name: "Iga Swiatek", Region: "POL", Age: 22, TournamentsPlayed: 19, Points: 9770},
	{Rank: 14, Name: "Daria Kasatkina", Region: "RUS", Age: 26, TournamentsPlayed: 25, Points: 2838},
	{Rank: 19, Name: "Elina Svitolina", Region: "UKR", Age: 29, TournamentsPlayed: 15, Points: 2212},
	{Rank: 8, Name: "Marketa Vondrousova", Region: "CZE", Age: 24, TournamentsPlayed: 16, Points: 3846},
	{Rank: 16, Name: "Veronika Kudermetova", Region: "RUS", Age: 26, TournamentsPlayed: 25, Points: 2495},
	{Rank: 6, Name: "Ons Jabeur", Region: "TUN", Age: 29, TournamentsPlayed: 20, Points: 4076},
	{Rank: 18, Name: "Petra Kvitova", Region: "CZE", Age: 33, TournamentsPlayed: 15, Points: 2465},
	{Rank: 13, Name: "Beatriz Haddad Maia", Region: "BRA", Age: 27, TournamentsPlayed: 23, Points: 2950},
	{Rank: 3, Name: "Coco Gauff", Region: "USA", Age: 19, TournamentsPlayed: 19, Points: 7200},
	{Rank: 11, Name: "Barbora Krejcikova", Region: "CZE", Age: 28, TournamentsPlayed: 21, Points: 3081},
	{Rank: 15, Name: "Liudmila Samsonova", Region: "RUS", Age: 25, TournamentsPlayed: 24, Points: 2700},
	{Rank: 20, Name: "Caroline Garcia", Region: "FRA", Age: 30, TournamentsPlayed: 27, Points: 2160},
	{Rank: 17, Name: "Madison Keys", Region: "USA", Age: 28, TournamentsPlayed: 17, Points: 2478},
	{Rank: 12, Name: "Jelena Ostapenko", Region: "LAT", Age: 26, TournamentsPlayed: 23, Points: 3028},
	{Rank: 9, Name: "Maria Sakkari", Region: "GRE", Age: 28, TournamentsPlayed: 24, Points: 3710},
	{Rank: 2, Name: "Aryna Sabalenka", Region: "BLR", Age: 25, TournamentsPlayed: 16, Points: 8905},
	{Rank: 7, Name: "Qinwen Zheng", Region: "CHN", Age: 21, TournamentsPlayed: 22, Points: 3950},
}

func ListPlayers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(players)
}

func GetPlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	playerRank := params["rank"]

	// Convert playerRank to integer
	rank, err := strconv.Atoi(playerRank)
	if err != nil {
		http.Error(w, "Invalid player rank", http.StatusBadRequest)
		return
	}

	for _, player := range players {
		if player.Rank == rank {
			json.NewEncoder(w).Encode(player)
			return
		}
	}

	http.NotFound(w, r)
}

func ListTopPlayers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Sort players by points in descending order
	sort.Slice(players, func(i, j int) bool {
		return players[i].Points > players[j].Points
	})

	// Take the top 20 players
	topPlayers := players[:20]

	json.NewEncoder(w).Encode(topPlayers)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("This app lists Top-20 women tennis players according to WTA! - Author: Daniyal Tuzelbayev"))
}

func SetupRoutes(r *mux.Router) {
	r.HandleFunc("/players", ListPlayers).Methods("GET")
	r.HandleFunc("/players/{rank}", GetPlayer).Methods("GET")
	r.HandleFunc("/top-players", ListTopPlayers).Methods("GET")
	r.HandleFunc("/health", HealthCheck).Methods("GET")
}
