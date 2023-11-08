package pubg

import (
	"encoding/json"
	"fmt"
)

// Clan objects contain information about a clan

// Clan structure.
type Clan struct {
	Data struct {
		// Identifier for this object type ("clan")
		Type string `json:"type"`
		// Clan ID
		ID         string `json:"id"`
		Attributes struct {
			ClanName        string `json:"clanName"`
			ClanTag         string `json:"clanTag"`
			ClanLevel       int    `json:"clanLevel"`
			ClanMemberCount int    `json:"clanMemberCount"`
		} `json:"attributes"`
	} `json:"data"`
	Links links `json:"links"`
}

// Clans returns information about the clan.
func (c Client) Clans(platform Platform, clanID string) (clan *Clan, err error) {
	b, _, err := c.requestGET(platform, fmt.Sprintf("/players/%s/clans", clanID))
	if err != nil {
		return
	}

	clan = &Clan{}
	err = json.Unmarshal(b, clan)
	return
}
