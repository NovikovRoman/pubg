package pubg

type links struct {
	// Link to this object
	Self string `json:"self"`
}

type dataIDAndType struct {
	Data struct {
		Type string `json:"type"`
		ID   string `json:"id"`
	} `json:"data"`
}
type dataArrayIDAndType struct {
	Data []struct {
		Type string `json:"type"`
		ID   string `json:"id"`
	} `json:"data"`
}

type statistics struct {
	Type       string `json:"type"`
	Attributes struct {
		GameModeStats map[string]gameModeStats `json:"gameModeStats"`
		BestRankPoint float64                  `json:"bestRankPoint"`
	} `json:"attributes"`

	Relationships struct {
		// type - Identifier for this object type ("player")
		// id - Player ID
		Player dataIDAndType `json:"player"`
		// type - Identifier for this object type ("season")
		// id - Season ID. Used to lookup a player's stats for this season on the /players endpoint
		Season dataIDAndType `json:"season"`

		// Contains a list of match IDs.
		// Array:
		//     type - Identifier for this object type ("match")
		//     id - Match ID. Used to lookup the full match object on the /matches endpoint
		MatchesSolo     dataArrayIDAndType `json:"matchesSolo"`
		MatchesSoloFPP  dataArrayIDAndType `json:"matchesSoloFPP"`
		MatchesDuo      dataArrayIDAndType `json:"matchesDuo"`
		MatchesDuoFPP   dataArrayIDAndType `json:"matchesDuoFPP"`
		MatchesSquad    dataArrayIDAndType `json:"matchesSquad"`
		MatchesSquadFPP dataArrayIDAndType `json:"matchesSquadFPP"`
	} `json:"relationships"`
}

type gameModeStats struct {
	// Number of enemy players this player damaged that were killed by teammates
	Assists int `json:"assists"`
	// Number of boost items used
	Boosts int `json:"boosts"`
	// Number of enemy players knocked
	DBNOs int `json:"dBNOs"`
	// Number of kills during the most recent day played.
	DailyKills int `json:"dailyKills"`
	// Total damage dealt. Note: Self inflicted damage is subtracted
	DamageDealt float64 `json:"damageDealt"`
	// Days :)
	Days int `json:"days"`
	// Number of wins during the most recent day played.
	DailyWins int `json:"dailyWins"`
	// Number of enemy players killed with headshots
	HeadshotKills int `json:"headshotKills"`
	// Number of healing items used
	Heals int `json:"heals"`
	// Number of enemy players killed
	Kills int `json:"kills"`
	// Longest kill
	LongestKill float64 `json:"longestKill"`
	// Longest time survived in a match
	LongestTimeSurvived float64 `json:"longestTimeSurvived"`
	// Number of matches lost
	Losses int `json:"losses"`
	// Maximum kill streaks
	MaxKillStreaks int `json:"maxKillStreaks"`
	// Longest time survived in a match
	MostSurvivalTime float64 `json:"mostSurvivalTime"`
	// Number of rank points the player was awarded. This value will be 0 when roundsPlayed < 10
	RankPoints float64 `json:"rankPoints"`
	// Rank title in the form title-level
	RankPointsTitle string `json:"rankPointsTitle"`
	// Number of times this player revived teammates
	Revives int `json:"revives"`
	// Total distance traveled in vehicles measured in meters
	RideDistance float64 `json:"rideDistance"`
	// Number of kills while in a vehicle
	RoadKills int `json:"roadKills"`
	// Highest number of kills in a single match
	RoundMostKills int `json:"roundMostKills"`
	// Number of matches played
	RoundsPlayed int `json:"roundsPlayed"`
	// Number of self-inflicted deaths
	Suicides int `json:"suicides"`
	// Total distance traveled while swimming measured in meters
	SwimDistance float64 `json:"swimDistance"`
	// Number of times this player killed a teammate
	TeamKills int `json:"teamKills"`
	// Total time survived
	TimeSurvived float64 `json:"timeSurvived"`
	// Number of times this player made it to the top 10 in a match
	Top10s int `json:"top10s"`
	// Number of vehicles destroyed
	VehicleDestroys int `json:"vehicleDestroys"`
	// Total distance traveled on foot measured in meters
	WalkDistance float64 `json:"walkDistance"`
	// Number of weapons picked up
	WeaponsAcquired int `json:"weaponsAcquired"`
	// Number of kills during the most recent week played
	WeeklyKills int `json:"weeklyKills"`
	// Number of wins during the most recent week played.
	WeeklyWins int `json:"weeklyWins"`
	// Number of matches won
	Wins int `json:"wins"`
}

// The weapon summary for each weapon
type weaponSummary struct {
	// The total amount of XP earned for this weapon
	XPTotal int `json:"XPTotal"`
	// The current level of this weapon
	LevelCurrent int `json:"LevelCurrent"`
	// The current tier of this weapon
	TierCurrent int `json:"TierCurrent"`
	// The weapon mastery stats for this weapon
	StatsTotal struct {
		// Most defeats in a single match
		MostDefeatsInAGame int `json:"MostDefeatsInAGame"`
		// The total number of defeats in their career
		Defeats int `json:"Defeats"`
		// The most damage that the player did in a single match
		MostDamagePlayerInAGame float64 `json:"MostDamagePlayerInAGame"`
		// The total damage that the player has done in their career
		DamagePlayer float64 `json:"DamagePlayer"`
		// The most headshots that the player did in a single match
		MostHeadShotsInAGame int `json:"MostHeadShotsInAGame"`
		// The total headshots that the player has done in their career
		HeadShots int `json:"HeadShots"`
		// The longest distance that the player got a defeat for
		LongestDefeat float64 `json:"LongestDefeat"`
		// The number of long range defeats for the player
		LongRangeDefeats int `json:"LongRangeDefeats"`
		// The total number of kills for the player
		Kills int `json:"Kills"`
		// The most kills for a player in a single match
		MostKillsInAGame int `json:"MostKillsInAGame"`
		// The total number of times that the player has caused another player to become groggy during their career
		Groggies int `json:"Groggies"`
		// The highest number of times that the player has caused another player to become groggy during a match
		MostGroggiesInAGame int `json:"MostGroggiesInAGame"`
	} `json:"StatsTotal"`

	// All of the medals received for this weapon
	Medals []struct {
		// The name of the medal
		MedalID string `json:"MedalId"`
		// The number of times that the player received the medal
		Count int `json:"Count"`
	} `json:"Medals"`
}

type rankedStatistics struct {
	Type       string `json:"type"`
	Attributes struct {
		RankedGameModeStats map[string]rankedGameModeStats `json:"rankedGameModeStats"`
	} `json:"attributes"`

	Relationships struct {
		// type - Identifier for this object type ("player")
		// id - Player ID
		Player dataIDAndType `json:"player"`
		// type - Identifier for this object type ("season")
		// id - Season ID. Used to lookup a player's stats for this season on the /players endpoint
		Season dataIDAndType `json:"season"`
	} `json:"relationships"`
}

// rankedGameModeStats structure contain a player's aggregated ranked stats for a game mode in the context of a season.
type rankedGameModeStats struct {
	// Player's current rank points
	CurrentRankPoint int `json:"currentRankPoint"`
	// Player's highest rank points
	BestRankPoint int `json:"bestRankPoint"`
	CurrentTier   struct {
		// Player's current ranked tier
		Tier string `json:"tier"`
		// Player's current ranked subtier
		SubTier string `json:"subTier"`
	} `json:"currentTier"`
	BestTier struct {
		// Player's current ranked tier
		Tier string `json:"tier"`
		// Player's current ranked subtier
		SubTier string `json:"subTier"`
	} `json:"bestTier"`
	// Number of matches played
	RoundsPlayed int `json:"roundsPlayed"`
	// Average rank
	AvgRank float64 `json:"avgRank"`
	// Ratio of number of times this player made it to the top 10 in a match / times didn't make it to top 10
	Top10Ratio float64 `json:"top10Ratio"`
	// Ratio of number of matches won / matches didn't win
	WinRatio float64 `json:"winRatio"`
	// Number of enemy players this player damaged that were killed by teammates
	Assists int `json:"assists"`
	// Number of matches won
	Wins int `json:"wins"`
	// Kill death assist ratio
	Kda float64 `json:"kda"`
	// Number of enemy players killed
	Kills int `json:"kills"`
	// Number of player deaths
	Deaths int `json:"deaths"`
	// Total damage dealt. Note: Self inflicted damage is subtracted
	DamageDealt float64 `json:"damageDealt"`
	// Number of enemy players knocked
	DBNOs int `json:"dBNOs"`
}
