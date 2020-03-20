package pubg

import (
	"encoding/json"
	"fmt"
)

// Weapon Mastery contains weapon summaries for the lifetime of a player.
type WeaponMastery struct {
	Data struct {
		// Identifier for this object type ("weaponMasterySummary")
		Type string `json:"type"`
		// Player ID (also known as account ID)
		ID         string `json:"id"`
		Attributes struct {
			// The platform
			Platform string `json:"platform"`
			// SeasonId: career
			SeasonID string `json:"seasonId"`
			// The match ID of the last completed match that was played.
			LatestMatchID string `json:"latestMatchId"`
			// The weapon summary for each weapon
			WeaponSummaries map[string]weaponSummary `json:"weaponSummaries"`
		} `json:"attributes"`
	} `json:"data"`

	Links links `json:"links"`
}

// WeaponMastery returns a weapon mastery information for a single player.
func (c Client) WeaponMastery(platform Platform, accountID string) (weaponMastery *WeaponMastery, err error) {
	b, _, err := c.requestGET(platform, fmt.Sprintf("/players/%s/weapon_mastery", accountID))
	if err != nil {
		return
	}

	weaponMastery = &WeaponMastery{}
	err = json.Unmarshal(b, weaponMastery)
	return
}
