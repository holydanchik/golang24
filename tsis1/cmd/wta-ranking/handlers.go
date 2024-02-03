package main

import (
	"encoding/json"
	"net/http"
	"sort"
	"strconv"

	"github.com/gorilla/mux"
	wtaranking "github.com/holydanchik/golang24/tsis1/pkg/wta-ranking"
)

func ListPlayers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(wtaranking.GetPlayers())
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

	for _, player := range wtaranking.GetPlayers() {
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
	sort.Slice(wtaranking.GetPlayers(), func(i, j int) bool {
		return wtaranking.GetPlayers()[i].Points > wtaranking.GetPlayers()[j].Points
	})

	// Take the top 20 players
	topPlayers := wtaranking.GetPlayers()[:20]

	json.NewEncoder(w).Encode(topPlayers)
}

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte(wtaranking.Info()))
}
