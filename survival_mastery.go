package pubg

import (
	"encoding/json"
	"fmt"
)

// Survival Mastery contains survival mastery data for a player

// SurvivalMastery structure.
type SurvivalMastery struct {
	Data struct {
		// Identifier for this object type ("survivalMasterySummary")
		Type string `json:"type"`
		// Player ID (also known as account ID)
		ID         string `json:"id"`
		Attributes struct {
			// Survival Mastery experience points
			XP int `json:"xp"`
			// Survival Mastery tier
			Tier int `json:"tier"`
			// Survival Mastery level
			Level int `json:"level"`
			// Number of matches played that count toward Survival Mastery
			TotalMatchesPlayed int `json:"totalMatchesPlayed"`
			// The match ID of the last completed match that was played
			LatestMatchId string `json:"latestMatchId"`
			Stats         struct {
				// Number of air drops called
				AirDropsCalled survivalMasteryStats `json:"airDropsCalled"`
				// Total amount of damage dealt to other players
				DamageDealt survivalMasteryStats `json:"damageDealt"`
				// Total amount of damage taken
				DamageTaken survivalMasteryStats `json:"damageTaken"`
				// Total distance travelled by swimming
				DistanceBySwimming survivalMasteryStats `json:"distanceBySwimming"`
				// Total distance travelled by vehicle
				DistanceByVehicle survivalMasteryStats `json:"distanceByVehicle"`
				// Total distance travelled on foot
				DistanceOnFoot survivalMasteryStats `json:"distanceOnFoot"`
				// Total distance travelled by foot, swimming, and vehicle
				DistanceTotal survivalMasteryStats `json:"distanceTotal"`
				// Total amount healed
				Healed survivalMasteryStats `json:"healed"`
				// Number of times landing in a hot drop location
				HotDropLandings survivalMasteryStats `json:"hotDropLandings"`
				// Number of enemy crates looted
				EnemyCratesLooted survivalMasteryStats `json:"enemyCratesLooted"`
				// Match placement
				Position survivalMasteryStats `json:"position"`
				// Number of times revived
				Revived survivalMasteryStats `json:"revived"`
				// Number of times reviving another player
				TeammatesRevived survivalMasteryStats `json:"teammatesRevived"`
				// Total time survived
				TimeSurvived survivalMasteryStats `json:"timeSurvived"`
				// Number of throwables thrown
				ThrowablesThrown survivalMasteryStats `json:"throwablesThrown"`
				// Number of times placing in the top 10
				Top10 survivalMasteryStats `json:"top10"`
			} `json:"stats"`
		} `json:"attributes"`
	} `json:"data"`
	Links links `json:"links"`
}

// Number of air drops called
type survivalMasteryStats struct {
	// Career total
	Total float64 `json:"total"`
	// Career average
	Average float64 `json:"average"`
	// Career best
	CareerBest float64 `json:"careerBest"`
	// Value in last match
	LastMatchValue float64 `json:"lastMatchValue"`
}

// SurvivalMastery returns a survival mastery information for a single player.
func (c Client) SurvivalMastery(platform Platform, accountID string) (survivalMastery *SurvivalMastery, err error) {
	b, _, err := c.requestGET(platform, fmt.Sprintf("/players/%s/survival_mastery", accountID))
	if err != nil {
		return
	}

	survivalMastery = &SurvivalMastery{}
	err = json.Unmarshal(b, survivalMastery)
	return
}
