package pubg

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// LifetimeStatsPlayer structure.
type LifetimeStatsPlayer struct {
	Data  statistics `json:"data"`
	Links links      `json:"links"`
}
// LifetimeStatsPlayers structure.
type LifetimeStatsPlayers struct {
	Data  []statistics `json:"data"`
	Links links        `json:"links"`
}

// LifetimeStatsPlayer returns lifetime stats for a single player.
func (c Client) LifetimeStatsPlayer(platform Platform, accountID string) (stats *LifetimeStatsPlayer, err error) {
	b, _, err := c.requestGET(platform, fmt.Sprintf("/players/%s/seasons/lifetime", accountID))
	if err != nil {
		return
	}

	stats = &LifetimeStatsPlayer{}
	err = json.Unmarshal(b, stats)
	return
}

// LifetimeStatsPlayers returns lifetime stats for up to 10 players.
func (c Client) LifetimeStatsPlayers(platform Platform, gameMode GameMode, accountID ...string) (stats *LifetimeStatsPlayers, err error) {
	if !gameMode.IsValid() {
		err = errors.New("Unknown game mode. ")
		return
	}

	if len(accountID) == 0 || len(accountID) > 10 {
		err = errors.New("You must specify 1 to 10 players. ")
		return
	}

	b, _, err := c.requestGET(platform,
		fmt.Sprintf("/seasons/lifetime/gameMode/%s/players?filter[playerIds]=%s", gameMode, strings.Join(accountID, ",")))
	if err != nil {
		return
	}

	stats = &LifetimeStatsPlayers{}
	err = json.Unmarshal(b, stats)
	return
}
