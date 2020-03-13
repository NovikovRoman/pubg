package pubg

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Leaderboards are updated every 2 hours
type Leaderboard struct {
	Data struct {
		// Identifier for this object type ("leaderboard")
		Type string `json:"type"`
		// A randomly generated ID assigned to this resource object for linking elsewhere in the leaderboard response
		ID         string `json:"id"`
		Attributes struct {
			// Platform shard
			ShardId string `json:"shardId"`
			// The game mode
			GameMode string `json:"gameMode"`
			// Season ID postfix
			SeasonId string `json:"seasonId"`
		} `json:"attributes"`

		Relationships struct {
			// Array:
			//     type - Identifier for this object type ("player")
			//     id - The accountId of the player
			Players dataArrayIDAndType `json:"players"`
		} `json:"relationships"`
	} `json:"data"`

	Included []struct {
		// Identifier for this object type ("player")
		Type string `json:"type"`
		// The accountId of the player
		ID string `json:"id"`

		Attributes struct {
			// The player's IGN
			Name string `json:"name"`
			// The player's current rank
			Rank int `json:"rank"`
			// Player stats in the context of a season
			Stats struct {
				// Number of rank points the player was awarded.
				RankPoints float64 `json:"rankPoints"`
				// Number of matches won
				Wins int `json:"wins"`
				// Number of games played
				Games int `json:"games"`
				// Win ratio
				WinRatio float64 `json:"winRatio"`
				// Average amount of damage dealt per match
				AverageDamage int `json:"averageDamage"`
				// Number of enemy players killed
				Kills int `json:"kills"`
				// Kill death ratio
				KillDeathRatio float64 `json:"killDeathRatio"`
				// Average rank
				AverageRank float64 `json:"averageRank"`
			} `json:"stats"`
		} `json:"attributes"`
	} `json:"included"`

	Links links `json:"links"`
}

// Get the leaderboard for a game mode.
func (c Client) Leaderboards(platform Platform, seasonID string, gameMode GameMode, page int) (leaderboard *Leaderboard, err error) {
	if !gameMode.IsValid() {
		err = errors.New("Unknown game mode. ")
		return
	}

	page--
	if page < 0 {
		page = 0
	}

	b, err := c.requestGET(platform, fmt.Sprintf("/leaderboards/%s/%s?page[number]=%d", seasonID, gameMode, page))
	if err != nil {
		return
	}

	leaderboard = &Leaderboard{}
	err = json.Unmarshal(b, leaderboard)
	return
}
