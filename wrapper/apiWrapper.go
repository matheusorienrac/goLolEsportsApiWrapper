package wrapper

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/matheusorienrac/goLolEsportsApiWrapper/enums"
	"github.com/matheusorienrac/goLolEsportsApiWrapper/models"
)

var client = &http.Client{}

// Sends a request to the specified endpoint with the specified query parameters.
func RequestLoLesportsAPI(endpoint string, queryParameters map[string]string) ([]byte, error) {
	req, err := http.NewRequest("GET", endpoint, nil)
	if err != nil {
		fmt.Println("there was an error creating the request: " + err.Error())
		return nil, err
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
		fmt.Println("there was an error sending the request: " + err.Error())
		return nil, err
	}

	bodyBytes, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("there was an error reading the response body: " + err.Error())
		return nil, err
	}

	return bodyBytes, nil
}

// Gets all games that are currently live.
func GetLiveGames(hl enums.HlType) (*models.Live, error) {
	getLiveGamesEndpoint := "https://esports-api.lolesports.com/persisted/gw/getLive"

	bodyBytes, err := RequestLoLesportsAPI(getLiveGamesEndpoint, map[string]string{"hl": hl.String()})
	if err != nil {
		fmt.Println("there was an error getting the live games: " + err.Error())
		return nil, err
	}

	fmt.Println(string(bodyBytes))

	var liveGames = &models.Live{}
	err = json.Unmarshal(bodyBytes, liveGames)
	if err != nil {
		fmt.Println("there was an error unmarshalling the liveGames: " + err.Error())
		return nil, err
	}

	return liveGames, nil
}

// Gets a window of match details for a game that is either live or already finished.
func GetWindow(gameID int64, startingTime time.Time) (*models.Window, error) {
	getWindowEndpoint := fmt.Sprintf("https://feed.lolesports.com/livestats/v1/window/%v", gameID)

	startingTime = transformTime(startingTime)
	bodyBytes, err := RequestLoLesportsAPI(getWindowEndpoint, map[string]string{"startingTime": startingTime.Format(time.RFC3339), "hl": enums.EnUS.String()})
	if err != nil {
		fmt.Println("there was an error getting the window: " + err.Error())
		return nil, err
	}

	fmt.Println(string(bodyBytes))
	var gameWindow = &models.Window{}
	err = json.Unmarshal(bodyBytes, gameWindow)
	if err != nil {
		fmt.Println("there was an error unmarshalling the gameWindow: " + err.Error())
		return nil, err
	}

	return gameWindow, nil
}

// Gets details of a game. participantIDs is a list of participant ids separated by underscores.
func GetDetails(gameID int64, startingTime time.Time, participantIDs string) (*models.Details, error) {

	getDetailsEndpoint := fmt.Sprintf("https://feed.lolesports.com/livestats/v1/details/%v", gameID)

	bodyBytes, err := RequestLoLesportsAPI(getDetailsEndpoint, map[string]string{"startingTime": startingTime.String(), "participantIds": participantIDs})
	if err != nil {
		fmt.Println("there was an error getting the details: " + err.Error())
		return nil, err
	}

	var gameDetails = &models.Details{}
	err = json.Unmarshal(bodyBytes, gameDetails)
	if err != nil {
		fmt.Println("there was an error unmarshalling the gameDetails: " + err.Error())
		return nil, err
	}

	return gameDetails, nil
}

// Transforms the time into a multiple of 10 and also removes 60 seconds so we can guarantee that
// the window we are getting is at least 45 seconds old.
func transformTime(startingTime time.Time) time.Time {
	// Truncate milliseconds
	startingTime = startingTime.Truncate(time.Second)

	// End time needs to be at least 45 seconds old.
	startingTime = startingTime.Add(-time.Duration(60) * time.Second)

	// Truncate seconds to be divisible by 10
	seconds := startingTime.Second()
	remainder := seconds % 10
	startingTime = startingTime.Add(-time.Duration(remainder) * time.Second)

	return startingTime
}
