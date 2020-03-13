package pubg

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

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

type SeasonStatsPlayer struct {
	Data  statistics `json:"data"`
	Links links      `json:"links"`
}

type SeasonStatsPlayers struct {
	Data  []statistics `json:"data"`
	Links links        `json:"links"`
}

// Get the list of available seasons.
func (c Client) Seasons(platform Platform) (seasons *Seasons, err error) {
	b, err := c.requestGET(platform, "/seasons")
	if err != nil {
		return
	}

	seasons = &Seasons{}
	err = json.Unmarshal(b, seasons)
	return
}

// Get season information for a single player.
func (c Client) SeasonStatsPlayer(platform Platform, seasonID, accountID string) (stats *SeasonStatsPlayer, err error) {
	b, err := c.requestGET(platform, fmt.Sprintf("/players/%s/seasons/%s", accountID, seasonID))
	if err != nil {
		return
	}

	stats = &SeasonStatsPlayer{}
	err = json.Unmarshal(b, stats)
	return
}

// Get season information for up to 10 players.
func (c Client) SeasonStatsPlayers(platform Platform, seasonID string, gameMode GameMode, accountID ...string) (stats *SeasonStatsPlayers, err error) {
	if !gameMode.IsValid() {
		err = errors.New("Unknown game mode. ")
		return
	}

	if len(accountID) == 0 || len(accountID) > 10 {
		err = errors.New("You must specify 1 to 10 players. ")
		return
	}

	b, err := c.requestGET(platform,
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
