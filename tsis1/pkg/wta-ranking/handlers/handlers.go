package handlers

import (
	"encoding/json"
	"net/http"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
	models "github.com/holydanchik/golang24/tsis1/pkg/wta-ranking/models"
	info "github.com/holydanchik/golang24/tsis1/pkg/wta-ranking/info"
)

func ListPlayers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(models.GetPlayers())
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

	for _, player := range models.GetPlayers() {
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
	sort.Slice(models.GetPlayers(), func(i, j int) bool {
		return models.GetPlayers()[i].Points > models.GetPlayers()[j].Points
	})

	// Take the top 20 players
	topPlayers := models.GetPlayers()[:20]

	json.NewEncoder(w).Encode(topPlayers)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(info.Info()))
}
