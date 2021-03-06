package pubg

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// Notes:
//
// - Use the platform shard when making requests for PC and PS4 players' season stats for seasons after
// division.bro.official.2018-09, and for Xbox season stats for seasons after division.bro.official.2018-08.
// Use the platform-region shard for making any other requests for players' season stats.
//
// - The list of seasons can be queried using both the plaform shard and the platform-region shard.

// Seasons structure.
type Seasons struct {
	Data []struct {
		// Identifier for this object type ("season")
		Type string `json:"type"`
		// Season ID. Used to lookup a player's stats for this season on the /players endpoint
		ID         string `json:"id"`
		Attributes struct {
			// Indicates if the season is active
			IsCurrentSeason bool `json:"isCurrentSeason"`
			// Indicates if the season is not active
			IsOffseason bool `json:"isOffseason"`
		} `json:"attributes"`
	} `json:"data"`

	Links links `json:"links"`
}

// RankedStatsPlayer structure.
type RankedStatsPlayer struct {
	Data  rankedStatistics `json:"data"`
	Links links            `json:"links"`
}

// SeasonStatsPlayer structure.
type SeasonStatsPlayer struct {
	Data  statistics `json:"data"`
	Links links      `json:"links"`
}

// SeasonStatsPlayers structure.
type SeasonStatsPlayers struct {
	Data  []statistics `json:"data"`
	Links links        `json:"links"`
}

// Seasons returns the list of available seasons.
func (c Client) Seasons(platform Platform) (seasons *Seasons, err error) {
	var b []byte
	if b, _, err = c.requestGET(platform, "/seasons"); err != nil {
		return
	}

	seasons = &Seasons{}
	err = json.Unmarshal(b, seasons)
	return
}

// RankedStats returns ranked stats for a single player.
func (c Client) RankedStatsPlayer(platform Platform, seasonID, accountID string) (stats *RankedStatsPlayer, err error) {
	var b []byte
	b, _, err = c.requestGET(platform, fmt.Sprintf("/players/%s/seasons/%s/ranked", accountID, seasonID))
	if err != nil {
		return
	}

	stats = &RankedStatsPlayer{}
	err = json.Unmarshal(b, stats)
	return
}

// SeasonStatsPlayer returns a season information for a single player.
func (c Client) SeasonStatsPlayer(platform Platform, seasonID, accountID string) (stats *SeasonStatsPlayer, err error) {
	var b []byte
	b, _, err = c.requestGET(platform, fmt.Sprintf("/players/%s/seasons/%s", accountID, seasonID))
	if err != nil {
		return
	}

	stats = &SeasonStatsPlayer{}
	err = json.Unmarshal(b, stats)
	return
}

// SeasonStatsPlayers returns a season information for up to 10 players.
func (c Client) SeasonStatsPlayers(platform Platform, seasonID string, gameMode GameMode, accountID ...string) (stats *SeasonStatsPlayers, err error) {
	var b []byte
	if !gameMode.IsValid() {
		err = errors.New("Unknown game mode. ")
		return
	}

	if len(accountID) == 0 || len(accountID) > 10 {
		err = errors.New("You must specify 1 to 10 players. ")
		return
	}

	b, _, err = c.requestGET(platform,
		fmt.Sprintf("/seasons/%s/gameMode/%s/players?filter[playerIds]=%s",
			seasonID, gameMode, strings.Join(accountID, ","),
		),
	)
	if err != nil {
		return
	}

	stats = &SeasonStatsPlayers{}
	err = json.Unmarshal(b, stats)
	return
}
