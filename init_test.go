package pubg

import (
	"log"
	"os"
	"time"
)

// For testing
var (
	cTest *Client
	// Когда проходим все тесты, необходимо притормаживать, тк ограничение 10 запросов в минуту
	// Не касается Matches() и Status().
	pause            = time.Microsecond
	testMatchID      = ""
	testTelemetryURL = ""
	testSeasonID     = ""
)

func init() {
	apiKey := os.Getenv("APIKEY")
	if apiKey == "" {
		log.Fatal("Set the environment variable APIKEY before retrying.")
	}

	cTest = NewClient(apiKey, nil)

	// Когда проходим все тесты, необходимо притормаживать, тк ограничение 10 запросов в минуту
	if os.Getenv("PAUSE") != "" {
		pause = time.Second * 10
	}

	seasons, err := cTest.Seasons(SteamPlatform)
	if err != nil {
		log.Fatalln(err)
	}
	for _, s := range seasons.Data {
		if s.Attributes.IsCurrentSeason {
			testSeasonID = s.ID
		}
	}

	// получим MatchID
	s, err := cTest.SeasonStatsPlayer(SteamPlatform, testSeasonID, testAccountID)
	if err != nil {
		log.Fatalln(err)
	}

	if len(s.Data.Relationships.MatchesSolo.Data) > 0 {
		testMatchID = s.Data.Relationships.MatchesSolo.Data[0].ID

	} else if len(s.Data.Relationships.MatchesSoloFPP.Data) > 0 {
		testMatchID = s.Data.Relationships.MatchesSoloFPP.Data[0].ID

	} else if len(s.Data.Relationships.MatchesDuo.Data) > 0 {
		testMatchID = s.Data.Relationships.MatchesDuo.Data[0].ID

	} else if len(s.Data.Relationships.MatchesDuoFPP.Data) > 0 {
		testMatchID = s.Data.Relationships.MatchesDuoFPP.Data[0].ID

	} else if len(s.Data.Relationships.MatchesSquad.Data) > 0 {
		testMatchID = s.Data.Relationships.MatchesSquad.Data[0].ID

	} else if len(s.Data.Relationships.MatchesSquadFPP.Data) > 0 {
		testMatchID = s.Data.Relationships.MatchesSquadFPP.Data[0].ID

	} else {
		log.Fatalln("Matches not found")
	}

	// получим testTelemetryURL
	matches, err := cTest.Matches(SteamPlatform, testMatchID)
	if err != nil {
		log.Fatalln(err)
	}
	for _, i := range matches.Assets {
		if i.Attributes.Name == "telemetry" {
			testTelemetryURL = i.Attributes.URL
			break
		}
	}

	if testTelemetryURL == "" {
		log.Fatalln("Telemetry URL not found")
	}
}
