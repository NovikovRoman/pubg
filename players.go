package pubg

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

// Player objects contain information about a player and a list of their recent matches (up to 14 days old).
// Note: player objects are specific to platform shards.

// Player structure.
type Player struct {
	Data  playerData `json:"data"`
	Links links      `json:"links"`
}

// Players structure.
type Players struct {
	Data  []playerData `json:"data"`
	Links links        `json:"links"`
}

type playerData struct {
	// Identifier for this object type ("player")
	Type string `json:"type"`
	// Player ID
	ID         string `json:"id"`
	Attributes struct {
		// Version of the game
		PatchVersion string `json:"patchVersion"`
		// PUBG IGN
		Name string `json:"name"`
		// Identifies the studio and game
		TitleID string `json:"titleId"`
		// Platform shard
		ShardID string `json:"shardId"`
		// Innocent, TemporaryBan, PermanentBan
		BanType string `json:"banType"`
		// Clan ID
		ClanId string `json:"clanId"`
	} `json:"attributes"`

	// References to resource objects related to this player
	Relationships struct {
		// A list of match ids.
		// Array:
		//     type - Identifier for this object type ("match")
		//     id - Match ID. Used to lookup the full match object on the /matches endpoint
		Matches dataArrayIDAndType `json:"matches"`
	} `json:"relationships"`

	Links links `json:"links"`
}

// Player returns a single player.
func (c Client) Player(platform Platform, accountID string) (player *Player, err error) {
	b, _, err := c.requestGET(platform, fmt.Sprintf("/players/%s", accountID))
	if err != nil {
		return
	}

	player = &Player{}
	err = json.Unmarshal(b, player)
	return
}

// PlayersByIDs returns a collection of up to 10 players by IDs.
func (c Client) PlayersByIDs(platform Platform, accountID ...string) (players *Players, err error) {
	if len(accountID) == 0 || len(accountID) > 10 {
		err = errors.New("You must specify 1 to 10 players. ")
		return
	}

	b, _, err := c.requestGET(platform,
		fmt.Sprintf("/players?filter[playerIds]=%s", strings.Join(accountID, ",")))
	if err != nil {
		return
	}

	players = &Players{}
	err = json.Unmarshal(b, players)
	return
}

// PlayersByNames returns a collection of up to 10 players by names.
func (c Client) PlayersByNames(platform Platform, name ...string) (players *Players, err error) {
	if len(name) == 0 || len(name) > 10 {
		err = errors.New("You must specify 1 to 10 players. ")
		return
	}

	b, _, err := c.requestGET(platform, fmt.Sprintf("/players?filter[playerNames]=%s", strings.Join(name, ",")))
	if err != nil {
		return
	}

	players = &Players{}
	err = json.Unmarshal(b, players)
	return
}
