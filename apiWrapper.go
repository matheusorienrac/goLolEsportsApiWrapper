package gololesportsapiwrapper

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/matheusorienrac/goLolEsportsApiWrapper/enums"
	"github.com/matheusorienrac/goLolEsportsApiWrapper/models"
)

var client = &http.Client{}

// Sends a request to the specified endpoint with the specified query parameters.
func RequestLoLesportsAPI(endpoint string, queryParameters map[string]string) []byte {
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		log.Fatal(err)
	}

	q := req.URL.Query()
	for key, value := range queryParameters {
		q.Add(key, value)
	}
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

	return bodyBytes
}

// Gets all games that are currently live.
func GetLiveGames(hl enums.HlType) models.Live {
	getLiveGamesEndpoint := "https://esports-api.lolesports.com/persisted/gw/getLive"

	bodyBytes := RequestLoLesportsAPI(getLiveGamesEndpoint, map[string]string{"hl": hl.String()})

	fmt.Println(string(bodyBytes))

	var liveGames = models.Live{}
	err := json.Unmarshal(bodyBytes, &liveGames)
	if err != nil {
		log.Fatal(err)
	}

	return liveGames
}

// Gets a window of match details for a game that is either live or already finished.
func GetWindow(gameID int64, startingTime time.Time) models.Window {
	getWindowEndpoint := fmt.Sprintf("https://feed.lolesports.com/livestats/v1/window/{%s}", gameID)

	bodyBytes := RequestLoLesportsAPI(getWindowEndpoint, map[string]string{"startingTime": startingTime.String()})

	var gameWindow = models.Window{}
	err := json.Unmarshal(bodyBytes, &gameWindow)
	if err != nil {
		log.Fatal(err)
	}

	return gameWindow
}

// Gets details of a game. participantIDs is a list of participant ids separated by underscores.
func GetDetails(gameID int64, startingTime time.Time, participantIDs string) models.Details {

	getDetailsEndpoint := fmt.Sprintf("https://feed.lolesports.com/livestats/v1/details/{%s}", gameID)

	bodyBytes := RequestLoLesportsAPI(getDetailsEndpoint, map[string]string{"startingTime": startingTime.String(), "participantIds": participantIDs})

	var gameDetails = models.Details{}
	err := json.Unmarshal(bodyBytes, &gameDetails)
	if err != nil {
		log.Fatal(err)
	}

	return gameDetails
}
