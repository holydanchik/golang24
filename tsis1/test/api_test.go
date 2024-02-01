package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/holydanchik/golang24/tsis1/api"
)

func TestListPlayers(t *testing.T) {
	req, err := http.NewRequest("GET", "/players", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/players", api.ListPlayers).Methods("GET")

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var resultPlayers []api.Player
	err = json.Unmarshal(rr.Body.Bytes(), &resultPlayers)
	if err != nil {
		t.Fatal(err)
	}

	expectedPlayers := []api.Player{
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

	if len(resultPlayers) != len(expectedPlayers) {
		t.Errorf("unexpected number of players. Got %d, want %d", len(resultPlayers), len(expectedPlayers))
	}

	for i, player := range resultPlayers {
		if player != expectedPlayers[i] {
			t.Errorf("unexpected player details at index %d. Got %+v, want %+v", i, player, expectedPlayers[i])
		}
	}

}

func TestGetPlayer(t *testing.T) {
	testCases := []struct {
		Rank           int
		ExpectedPlayer api.Player
		ExpectedStatus int
	}{
		{
			Rank: 1,
			ExpectedPlayer: api.Player{
				Rank:              1,
				Name:              "Iga Swiatek",
				Region:            "POL",
				Age:               22,
				TournamentsPlayed: 19,
				Points:            9770,
			},
			ExpectedStatus: http.StatusOK,
		},
		{
			Rank: 5,
			ExpectedPlayer: api.Player{
				Rank:              5,
				Name:              "Elena Rybakina",
				Region:            "KAZ",
				Age:               24,
				TournamentsPlayed: 18,
				Points:            5688,
			},
			ExpectedStatus: http.StatusOK,
		},
	}

	for _, testCase := range testCases {
		t.Run(fmt.Sprintf("Rank %d", testCase.Rank), func(t *testing.T) {
			req, err := http.NewRequest("GET", fmt.Sprintf("/players/%d", testCase.Rank), nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			router := mux.NewRouter()
			router.HandleFunc("/players/{rank}", api.GetPlayer).Methods("GET")

			router.ServeHTTP(rr, req)

			if status := rr.Code; status != testCase.ExpectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, testCase.ExpectedStatus)
			}

			var resultPlayer api.Player
			err = json.Unmarshal(rr.Body.Bytes(), &resultPlayer)
			if err != nil {
				t.Fatal(err)
			}

			if resultPlayer != testCase.ExpectedPlayer {
				t.Errorf("unexpected player details. Got %+v, want %+v", resultPlayer, testCase.ExpectedPlayer)
			}
		})
	}
}

func TestListTopPlayers(t *testing.T) {
	req, err := http.NewRequest("GET", "/top-players", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	router := mux.NewRouter()
	router.HandleFunc("/top-players", api.ListTopPlayers).Methods("GET")

	router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var resultTopPlayers []api.Player
	err = json.Unmarshal(rr.Body.Bytes(), &resultTopPlayers)
	if err != nil {
		t.Fatal(err)
	}

	expectedTopPlayers := []api.Player{
		{Rank: 1, Name: "Iga Swiatek", Region: "POL", Age: 22, TournamentsPlayed: 19, Points: 9770},
		{Rank: 2, Name: "Aryna Sabalenka", Region: "BLR", Age: 25, TournamentsPlayed: 16, Points: 8905},
		{Rank: 3, Name: "Coco Gauff", Region: "USA", Age: 19, TournamentsPlayed: 19, Points: 7200},
		{Rank: 4, Name: "Jessica Pegula", Region: "USA", Age: 29, TournamentsPlayed: 21, Points: 5705},
		{Rank: 5, Name: "Elena Rybakina", Region: "KAZ", Age: 24, TournamentsPlayed: 18, Points: 5688},
		{Rank: 6, Name: "Ons Jabeur", Region: "TUN", Age: 29, TournamentsPlayed: 20, Points: 4076},
		{Rank: 7, Name: "Qinwen Zheng", Region: "CHN", Age: 21, TournamentsPlayed: 22, Points: 3950},
		{Rank: 8, Name: "Marketa Vondrousova", Region: "CZE", Age: 24, TournamentsPlayed: 16, Points: 3846},
		{Rank: 9, Name: "Maria Sakkari", Region: "GRE", Age: 28, TournamentsPlayed: 24, Points: 3710},
		{Rank: 10, Name: "Karolina Muchova", Region: "CZE", Age: 27, TournamentsPlayed: 14, Points: 3520},
		{Rank: 11, Name: "Barbora Krejcikova", Region: "CZE", Age: 28, TournamentsPlayed: 21, Points: 3081},
		{Rank: 12, Name: "Jelena Ostapenko", Region: "LAT", Age: 26, TournamentsPlayed: 23, Points: 3028},
		{Rank: 13, Name: "Beatriz Haddad Maia", Region: "BRA", Age: 27, TournamentsPlayed: 23, Points: 2950},
		{Rank: 14, Name: "Daria Kasatkina", Region: "RUS", Age: 26, TournamentsPlayed: 25, Points: 2838},
		{Rank: 15, Name: "Liudmila Samsonova", Region: "RUS", Age: 25, TournamentsPlayed: 24, Points: 2700},
		{Rank: 16, Name: "Veronika Kudermetova", Region: "RUS", Age: 26, TournamentsPlayed: 25, Points: 2495},
		{Rank: 17, Name: "Madison Keys", Region: "USA", Age: 28, TournamentsPlayed: 17, Points: 2478},
		{Rank: 18, Name: "Petra Kvitova", Region: "CZE", Age: 33, TournamentsPlayed: 15, Points: 2465},
		{Rank: 19, Name: "Elina Svitolina", Region: "UKR", Age: 29, TournamentsPlayed: 15, Points: 2212},
		{Rank: 20, Name: "Caroline Garcia", Region: "FRA", Age: 30, TournamentsPlayed: 27, Points: 2160},
	}

	if len(resultTopPlayers) != len(expectedTopPlayers) {
		t.Errorf("unexpected number of top players. Got %d, want %d", len(resultTopPlayers), len(expectedTopPlayers))
	}

	for i, player := range resultTopPlayers {
		if player != expectedTopPlayers[i] {
			t.Errorf("unexpected top player details at index %d. Got %+v, want %+v", i, player, expectedTopPlayers[i])
		}
	}

}
