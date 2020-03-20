package pubg

type telemetryObjectBlueZoneCustomOptions struct {
	PhaseNum                 int     `json:"phaseNum"`
	StartDelay               int     `json:"startDelay"`
	WarningDuration          int     `json:"warningDuration"`
	ReleaseDuration          int     `json:"releaseDuration"`
	PoisonGasDamagePerSecond float64 `json:"poisonGasDamagePerSecond"`
	RadiusRate               float64 `json:"radiusRate"`
	SpreadRatio              float64 `json:"spreadRatio"`
	LandRatio                float64 `json:"landRatio"`
	CircleAlgorithm          int     `json:"circleAlgorithm"`
}

type telemetryObjectCharacter struct {
	Name         string                  `json:"name"`
	TeamID       int                     `json:"teamId"`
	Health       float64                 `json:"health"`
	Location     telemetryObjectLocation `json:"location"`
	Ranking      int                     `json:"ranking"`
	AccountID    string                  `json:"accountId"`
	IsInBlueZone bool                    `json:"isInBlueZone"`
	IsInRedZone  bool                    `json:"isInRedZone"`
	Zone         []string                `json:"zone"`
}

// IsGame represents the phase of the game defined by the status of bluezone and safezone
// isGame = 0 -> Before lift off
//
// isGame = 0.1 -> On airplane
//
// isGame = 0.5 -> When there’s no ‘zone’ on map(before game starts)
//
// isGame = 1.0 -> First safezone and bluezone appear
//
// isGame = 1.5 -> First bluezone shrinks
//
// isGame = 2.0 -> Second bluezone appears
//
// isGame = 2.5 -> Second bluezone shrinks
//
// …
type telemetryObjectCommon struct {
	IsGame float64 `json:"isGame"`
}

type telemetryObjectGameResult struct {
	Rank       int                  `json:"rank"`
	GameResult string               `json:"gameResult"`
	TeamID     int                  `json:"teamId"`
	Stats      telemetryObjectStats `json:"stats"`
	AccountID  string               `json:"accountId"`
}

// Shows winning players only
type telemetryObjectGameResultOnFinished struct {
	Results []telemetryObjectGameResult `json:"results"`
}

type telemetryObjectGameState struct {
	ElapsedTime              int                     `json:"elapsedTime"`
	NumAliveTeams            int                     `json:"numAliveTeams"`
	NumJoinPlayers           int                     `json:"numJoinPlayers"`
	NumStartPlayers          int                     `json:"numStartPlayers"`
	NumAlivePlayers          int                     `json:"numAlivePlayers"`
	SafetyZonePosition       telemetryObjectLocation `json:"safetyZonePosition"`
	SafetyZoneRadius         float64                 `json:"safetyZoneRadius"`
	PoisonGasWarningPosition telemetryObjectLocation `json:"poisonGasWarningPosition"`
	PoisonGasWarningRadius   float64                 `json:"poisonGasWarningRadius"`
	RedZonePosition          telemetryObjectLocation `json:"redZonePosition"`
	RedZoneRadius            float64                 `json:"redZoneRadius"`
	BlackZonePosition        telemetryObjectLocation `json:"blackZonePosition"`
	BlackZoneRadius          float64                 `json:"blackZoneRadius"`
}

type telemetryObjectItem struct {
	ItemID        string   `json:"itemId"`
	StackCount    int      `json:"stackCount"`
	Category      string   `json:"category"`
	SubCategory   string   `json:"subCategory"`
	AttachedItems []string `json:"attachedItems"`
}

// ItemPackage structure.
type ItemPackage struct {
	ItemPackageID string                  `json:"itemPackageId"`
	Location      telemetryObjectLocation `json:"location"`
	Items         []telemetryObjectItem   `json:"items"`
}

//Location values are measured in centimeters. (0,0) is at the top-left of each map.
//
//The range for the X and Y axes is 0 - 816,000 for Erangel and Miramar.
//
//The range for the X and Y axes is 0 - 408,000 for Sanhok.
//
//The range for the X and Y axes is 0 - 612,000 for Vikendi.
//
//The range for the X and Y axes is 0 - 204,000 for Range.
type telemetryObjectLocation struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`
}

type telemetryObjectStats struct {
	KillCount           int     `json:"killCount"`
	DistanceOnFoot      float64 `json:"distanceOnFoot"`
	DistanceOnSwim      float64 `json:"distanceOnSwim"`
	DistanceOnVehicle   float64 `json:"distanceOnVehicle"`
	DistanceOnParachute float64 `json:"distanceOnParachute"`
	DistanceOnFreefall  float64 `json:"distanceOnFreefall"`
}

type telemetryObjectVehicle struct {
	VehicleType     string  `json:"vehicleType"`
	VehicleID       string  `json:"vehicleId"`
	VehicleUniqueID int     `json:"vehicleUniqueId"`
	HealthPercent   float64 `json:"healthPercent"`
	FeulPercent     float64 `json:"feulPercent"`
	SeatIndex       int     `json:"seatIndex"`
	IsWheelsInAir   bool    `json:"isWheelsInAir"`
	IsInWaterVolume bool    `json:"isInWaterVolume"`
}
