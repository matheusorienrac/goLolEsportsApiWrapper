package apiWrapper

import (
	"encoding/json"
	"fmt"
	"go-lol-esports-api-wrapper/enums"
	"go-lol-esports-api-wrapper/models"
	"io"
	"log"
	"net/http"
	"time"
)

var client = &http.Client{}

// Gets all games that are currently live.
func getLiveGames(hl enums.HlType) models.Live {
	req, err := http.NewRequest("GET", "https://esports-api.lolesports.com/persisted/gw/getLive", nil)
	if err != nil {
		log.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("hl", hl.String())
	req.URL.RawQuery = q.Encode()

	// Adds the required api key to the header
	req.Header.Add("x-api-key", "0TvQnueqKa5mxJntVWt0w4LpLfEkrV1Ta8rQBb9Z")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var liveGames = models.Live{}
	err = json.Unmarshal(bodyBytes, &liveGames)
	return liveGames
}

// Gets a window of match details for a game that is either live or already finished.
func getWindow(gameID int64, startingTime time.Time) models.Window {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://feed.lolesports.com/livestats/v1/window/{%s}", gameID), nil)
	if err != nil {
		log.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("startingTime", startingTime.String())
	req.URL.RawQuery = q.Encode()

	// Adds the required api key to the header
	req.Header.Add("x-api-key", "0TvQnueqKa5mxJntVWt0w4LpLfEkrV1Ta8rQBb9Z")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var gameWindow = models.Window{}
	err = json.Unmarshal(bodyBytes, &gameWindow)
	return gameWindow
}

// Gets details of a game. participantIDs is a list of participant ids separated by underscores.
func getDetails(gameID int64, startingTime time.Time, participantIDs string) models.Details {
	req, err := http.NewRequest("GET", fmt.Sprintf("https://feed.lolesports.com/livestats/v1/details/{%s}", gameID), nil)
	if err != nil {
		log.Fatal(err)
	}

	q := req.URL.Query()
	q.Add("startingTime", startingTime.String())
	q.Add("participantIds", participantIDs.String())

	req.URL.RawQuery = q.Encode()

	// Adds the required api key to the header
	req.Header.Add("x-api-key", "0TvQnueqKa5mxJntVWt0w4LpLfEkrV1Ta8rQBb9Z")

	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var gameDetails = models.Details{}
	err = json.Unmarshal(bodyBytes, &gameDetails)
	return gameDetails
}
