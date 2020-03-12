package pubg

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type LifetimeStatsPlayer struct {
	Data  statistics `json:"data"`
	Links links      `json:"links"`
}

type LifetimeStatsPlayers struct {
	Data  []statistics `json:"data"`
	Links links        `json:"links"`
}

// Get lifetime stats for a single player.
func (c Client) LifetimeStatsPlayer(platform Platform, accountID string) (stats *LifetimeStatsPlayer, err error) {
	b, err := c.requestGET(platform, fmt.Sprintf("/players/%s/seasons/lifetime", accountID))
	if err != nil {
		return
	}

	stats = &LifetimeStatsPlayer{}
	err = json.Unmarshal(b, stats)
	return
}

// Get lifetime stats for up to 10 players.
func (c Client) LifetimeStatsPlayers(platform Platform, gameMode GameMode, accountID ...string) (stats *LifetimeStatsPlayers, err error) {
	if len(accountID) == 0 || len(accountID) > 10 {
		err = errors.New("You must specify 1 to 10 players. ")
		return
	}

	b, err := c.requestGET(platform, fmt.Sprintf("/seasons/lifetime/gameMode/%s/players?filter[playerIds]=%s", gameMode, strings.Join(accountID, ",")))
	if err != nil {
		return
	}

	stats = &LifetimeStatsPlayers{}
	err = json.Unmarshal(b, stats)
	return
}
