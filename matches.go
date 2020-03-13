package pubg

import (
	"encoding/json"
	"fmt"
	"regexp"
	"time"
)

// Match objects contain information about a completed match such as the game mode played, duration, and which players participated.
type Matches struct {
	Data struct {
		// Identifier for this object type ("match")
		Type string `json:"type"`
		// Match ID
		ID         string `json:"id"`
		Attributes struct {
			// Time this match object was stored in the API
			CreatedAtRaw string    `json:"createdAt"`
			CreatedAt    time.Time `json:"-"`
			// Length of the match measured in seconds
			Duration int `json:"duration"`
			// Type of match [arcade, custom, event, official, training]
			MatchType string `json:"latestMatchId"`
			// Game mode played [duo, duo-fpp, solo, solo-fpp, squad, squad-fpp, conquest-duo, conquest-duo-fpp,
			//     conquest-solo, conquest-solo-fpp, conquest-squad, conquest-squad-fpp, esports-duo, esports-duo-fpp,
			//     esports-solo, esports-solo-fpp, esports-squad, esports-squad-fpp, normal-duo, normal-duo-fpp,
			//     normal-solo, normal-solo-fpp, normal-squad, normal-squad-fpp, tdm, war-duo, war-duo-fpp, war-solo,
			//     war-solo-fpp, war-squad, war-squad-fpp, zombie-duo, zombie-duo-fpp, zombie-solo, zombie-solo-fpp,
			//     zombie-squad, zombie-squad-fpp]
			GameMode string
			// Map name [Baltic_Main, Desert_Main, DihorOtok_Main, Erangel_Main, Range_Main, Savage_Main, Summerland_Main]
			MapName string `json:"mapName"`
			// True if this match is a custom match
			IsCustomMatch bool `json:"isCustomMatch"`
			// The state of the season [closed, prepare, progress]
			SeasonState string `json:"seasonState"`
			// Platform shard
			ShardId string `json:"shardId"`
			// Identifies the studio and game
			TitleId string `json:"titleId"`
		} `json:"attributes"`

		// References to resource objects that can be found in the included array
		Relationships struct {
			// Array:
			//     type - Identifier for this object type ("asset")
			//     id - AssetID. Used to find the full asset object in the included array
			Assets dataArrayIDAndType `json:"assets"`

			// Array:
			//     type - Identifier for this object type ("roster")
			//     id - RosterID. Used to find the full roster object in the included array
			Rosters dataArrayIDAndType `json:"rosters"`
		} `json:"relationships"`
	} `json:"data"`

	IncludedRaw  []json.RawMessage `json:"included"`
	Participants []participant     `json:"-"`

	// Rosters track the scores of each opposing group of participants. Rosters can have one or many participants
	// depending on the game mode. Roster objects are only meaningful within the context of a match and are not exposed
	// as a standalone resource.
	Rosters []roster `json:"-"`

	// Asset objects contain a URL string that links to a telemetry.json file, which will contain an array of event
	// objects that provide further insight into a match.
	Assets []asset `json:"-"`

	Links links `json:"links"`
}

type roster struct {
	// Identifier for this object type ("asset")
	Type string `json:"type"`
	// A randomly generated ID assigned to this resource object for linking elsewhere in the match response
	ID         string `json:"id"`
	Attributes struct {
		// Platform shard
		ShardId string `json:"shardId"`
		// Indicates if this roster won the match
		Won string `json:"won"`
		// Telemetry
		Stats struct {
			// This roster's placement in the match. 1 - 130
			Rank int `json:"rank"`
			// An arbitrary ID assigned to this roster
			TeamId int `json:"teamId"`
		} `json:"stats"`
	} `json:"attributes"`

	Relationships struct {
		// An array of references to participant objects that can be found in the included array
		// Array:
		//     type - Identifier for this object type ("participant")
		//     id - ParticipantID. Use to find full participant object in the included array
		Participants dataArrayIDAndType `json:"participants"`
	} `json:"relationships"`
}

type participant struct {
	// Identifier for this object type ("participant")
	Type string `json:"type"`
	// A randomly generated ID assigned to this resource object for linking elsewhere in the match response
	ID         string `json:"id"`
	Attributes struct {
		// Platform shard
		ShardId string `json:"shardId"`
		// Player stats in the context of a match
		Stats struct {
			// Number of players knocked
			DBNOs int `json:"DBNOs"`
			// Number of enemy players this player damaged that were killed by teammates. 0 - 128
			Assists int `json:"assists"`
			// Number of boost items used
			Boosts int `json:"boosts"`
			// Total damage dealt. Note: Self inflicted damage is subtracted
			DamageDealt float64 `json:"damageDealt"`
			// The way by which this player died, or alive if they didn't [alive, byplayer, byzone, suicide, logout]
			DeathType string `json:"deathType"`
			// Number of enemy players killed with headshots. 0 - 129
			HeadshotKills int `json:"headshotKills"`
			// Number of healing items used
			Heals int `json:"heals"`
			// This player's rank in the match based on kills. 1 - 130
			KillPlace int `json:"killPlace"`
			// Total number of kill streaks
			KillStreaks int `json:"killStreaks"`
			// Number of enemy players killed. 0 - 129
			Kills int `json:"kills"`
			// Longest kill
			LongestKill float64 `json:"longestKill"`
			// PUBG IGN of the player associated with this participant
			Name string `json:"name"`
			// Account ID of the player associated with this participant
			PlayerId string `json:"playerId"`
			// Number of times this player revived teammates
			Revives int `json:"revives"`
			// Total distance traveled in vehicles measured in meters
			RideDistance float64 `json:"rideDistance"`
			// Number of kills while in a vehicle
			RoadKills int `json:"roadKills"`
			// Total distance traveled while swimming measured in meters
			SwimDistance float64 `json:"swimDistance"`
			// Number of times this player killed a teammate
			TeamKills int `json:"teamKills"`
			// Amount of time survived measured in seconds
			TimeSurvived float64 `json:"timeSurvived"`
			// Number of vehicles destroyed
			VehicleDestroys int `json:"vehicleDestroys"`
			// Total distance traveled on foot measured in meters`
			WalkDistance float64 `json:"walkDistance"`
			// Number of weapons picked up
			WeaponsAcquired int `json:"weaponsAcquired"`
			// This player's placement in the match. 1 - 130
			WinPlace int `json:"winPlace"`
		} `json:"stats"`
	} `json:"attributes"`
}

type asset struct {
	// Identifier for this object type ("asset")
	Type string `json:"type"`
	// A randomly generated ID assigned to this resource object for linking elsewhere in the match response
	ID         string `json:"id"`
	Attributes struct {
		// Link to the telemetry.json file
		URL string `json:"URL"`
		// Time of telemetry creation
		CreatedAtRaw string    `json:"createdAt"`
		CreatedAt    time.Time `json:"-"`
		// Telemetry
		Name string `json:"name"`
	} `json:"attributes"`
}

// Get a single match.
func (c Client) Matches(platform Platform, matchID string) (matches *Matches, err error) {
	b, _, err := c.requestGET(platform, fmt.Sprintf("/matches/%s", matchID))
	if err != nil {
		return
	}

	matches = &Matches{}
	err = json.Unmarshal(b, matches)
	if err != nil {
		return
	}

	matches.Data.Attributes.CreatedAt, _ = time.Parse(layoutDateTime, matches.Data.Attributes.CreatedAtRaw)

	reRoster := regexp.MustCompile(`(?si)^{\s*"type"\s*:\s*"roster"`)
	reParticipant := regexp.MustCompile(`(?si)^{\s*"type"\s*:\s*"participant"`)
	reAsset := regexp.MustCompile(`(?si)^{\s*"type"\s*:\s*"asset"`)
	for _, raw := range matches.IncludedRaw {
		if reParticipant.Match(raw) {
			p := participant{}
			if err = json.Unmarshal(raw, &p); err != nil {
				return
			}
			matches.Participants = append(matches.Participants, p)

		} else if reRoster.Match(raw) {
			r := roster{}
			if err = json.Unmarshal(raw, &r); err != nil {
				return
			}
			matches.Rosters = append(matches.Rosters, r)

		} else if reAsset.Match(raw) {
			a := asset{}
			if err = json.Unmarshal(raw, &a); err != nil {
				return
			}
			a.Attributes.CreatedAt, _ = time.Parse(layoutDateTime, a.Attributes.CreatedAtRaw)
			matches.Assets = append(matches.Assets, a)
		}
	}

	return
}
