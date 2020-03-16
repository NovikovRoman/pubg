package pubg

import (
	"encoding/json"
	"time"
)

type telemetryEvent struct {
	DateRaw string    `json:"_D"`
	Type    string    `json:"_T"`
	Date    time.Time `json:"_"`
}

type LogArmorDestroy struct {
	telemetryEvent
	AttackId           int                      `json:"attackId"`
	Attacker           telemetryObjectCharacter `json:"attacker"`
	Victim             telemetryObjectCharacter `json:"victim"`
	DamageTypeCategory string                   `json:"damageTypeCategory"`
	DamageReason       string                   `json:"damageReason"`
	DamageCauserName   string                   `json:"damageCauserName"`
	Item               telemetryObjectItem      `json:"item"`
	Distance           float64                  `json:"distance"`
}

func NewLogArmorDestroy(raw json.RawMessage) (l *LogArmorDestroy, err error) {
	l = &LogArmorDestroy{}
	err = json.Unmarshal(raw, l)
	if err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

type LogBlackZoneEnded struct {
	telemetryEvent
	Survivors []telemetryObjectCharacter `json:"survivors"`
}

func NewLogBlackZoneEnded(raw json.RawMessage) (l *LogBlackZoneEnded, err error) {
	l = &LogBlackZoneEnded{}
	err = json.Unmarshal(raw, l)
	if err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

type LogCarePackageLand struct {
	telemetryEvent
	ItemPackage ItemPackage `json:"itemPackage"`
}

func NewLogCarePackageLand(raw json.RawMessage) (l *LogCarePackageLand, err error) {
	l = &LogCarePackageLand{}
	err = json.Unmarshal(raw, l)
	if err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

type LogCarePackageSpawn struct {
	telemetryEvent
	ItemPackage ItemPackage `json:"itemPackage"`
}

func NewLogCarePackageSpawn(raw json.RawMessage) (l *LogCarePackageSpawn, err error) {
	l = &LogCarePackageSpawn{}
	err = json.Unmarshal(raw, l)
	if err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

type LogGameStatePeriodic struct {
	telemetryEvent
	GameState telemetryObjectGameState `json:"gameState"`
}

func NewLogGameStatePeriodic(raw json.RawMessage) (l *LogGameStatePeriodic, err error) {
	l = &LogGameStatePeriodic{}
	err = json.Unmarshal(raw, l)
	if err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

type LogHeal struct {
	telemetryEvent
	Character  telemetryObjectCharacter `json:"character"`
	Item       telemetryObjectItem      `json:"item"`
	HealAmount float64                  `json:"healAmount"`
}

func NewLogHeal(raw json.RawMessage) (l *LogHeal, err error) {
	l = &LogHeal{}
	err = json.Unmarshal(raw, l)
	if err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

type LogItemAttach struct {
	telemetryEvent
	Character  telemetryObjectCharacter `json:"character"`
	ParentItem telemetryObjectItem      `json:"parentItem"`
	ChildItem  telemetryObjectItem      `json:"childItem"`
}

func NewLogItemAttach(raw json.RawMessage) (l *LogItemAttach, err error) {
	l = &LogItemAttach{}
	err = json.Unmarshal(raw, l)
	if err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

type LogItemDetach struct {
	telemetryEvent
	Character  telemetryObjectCharacter `json:"character"`
	ParentItem telemetryObjectItem      `json:"parentItem"`
	ChildItem  telemetryObjectItem      `json:"childItem"`
}

func NewLogItemDetach(raw json.RawMessage) (l *LogItemDetach, err error) {
	l = &LogItemDetach{}
	err = json.Unmarshal(raw, l)
	if err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

type LogItemDrop struct {
	telemetryEvent
	Character telemetryObjectCharacter `json:"character"`
	Item      telemetryObjectItem      `json:"item"`
}

func NewLogItemDrop(raw json.RawMessage) (l *LogItemDrop, err error) {
	l = &LogItemDrop{}
	err = json.Unmarshal(raw, l)
	if err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

type LogItemEquip struct {
	telemetryEvent
	Character telemetryObjectCharacter `json:"character"`
	Item      telemetryObjectItem      `json:"item"`
}

func NewLogItemEquip(raw json.RawMessage) (l *LogItemEquip, err error) {
	l = &LogItemEquip{}
	err = json.Unmarshal(raw, l)
	if err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

type LogItemPickup struct {
	telemetryEvent
	Character telemetryObjectCharacter `json:"character"`
	Item      telemetryObjectItem      `json:"item"`
}

func NewLogItemPickup(raw json.RawMessage) (l *LogItemPickup, err error) {
	l = &LogItemPickup{}
	err = json.Unmarshal(raw, l)
	if err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

type LogItemPickupFromCarepackage struct {
	telemetryEvent
	Character           telemetryObjectCharacter `json:"character"`
	Item                telemetryObjectItem      `json:"item"`
	CarePackageUniqueId float64                  `json:"carePackageUniqueId"`
}

func NewLogItemPickupFromCarepackage(raw json.RawMessage) (l *LogItemPickupFromCarepackage, err error) {
	l = &LogItemPickupFromCarepackage{}
	err = json.Unmarshal(raw, l)
	if err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

type LogItemPickupFromLootBox struct {
	telemetryEvent
	Character        telemetryObjectCharacter `json:"character"`
	Item             telemetryObjectItem      `json:"item"`
	OwnerTeamId      int                      `json:"ownerTeamId"`
	CreatorAccountId string                   `json:"creatorAccountId"`
}

func NewLogItemPickupFromLootBox(raw json.RawMessage) (l *LogItemPickupFromLootBox, err error) {
	l = &LogItemPickupFromLootBox{}
	err = json.Unmarshal(raw, l)
	if err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

type LogItemUnequip struct {
	telemetryEvent
	Character telemetryObjectCharacter `json:"character"`
	Item      telemetryObjectItem      `json:"item"`
}

func NewLogItemUnequip(raw json.RawMessage) (l *LogItemUnequip, err error) {
	l = &LogItemUnequip{}
	err = json.Unmarshal(raw, l)
	if err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

type LogItemUse struct {
	telemetryEvent
	Character telemetryObjectCharacter `json:"character"`
	Item      telemetryObjectItem      `json:"item"`
}

func NewLogItemUse(raw json.RawMessage) (l *LogItemUse, err error) {
	l = &LogItemUse{}
	err = json.Unmarshal(raw, l)
	if err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

type LogMatchDefinition struct {
	telemetryEvent
	MatchId     string `json:"MatchId"`
	PingQuality string `json:"PingQuality"`
	SeasonState string `json:"SeasonState"`
}

func NewLogMatchDefinition(raw json.RawMessage) (l *LogMatchDefinition, err error) {
	l = &LogMatchDefinition{}
	err = json.Unmarshal(raw, l)
	if err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

type LogMatchEnd struct {
	telemetryEvent
	Characters []telemetryObjectCharacter `json:"characters"`
	// Shows winning players only
	GameResultOnFinished telemetryObjectGameResultOnFinished `json:"gameResultOnFinished"`
}

func NewLogMatchEnd(raw json.RawMessage) (l *LogMatchEnd, err error) {
	l = &LogMatchEnd{}
	err = json.Unmarshal(raw, l)
	if err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

type LogMatchStart struct {
	telemetryEvent
	MapName                  string                                 `json:"mapName"`
	WeatherId                string                                 `json:"weatherId"`
	Characters               []telemetryObjectCharacter             `json:"characters"`
	CameraViewBehaviour      string                                 `json:"cameraViewBehaviour"`
	TeamSize                 int                                    `json:"teamSize"`
	IsCustomGame             bool                                   `json:"isCustomGame"`
	IsEventMode              bool                                   `json:"isEventMode"`
	BlueZoneCustomOptions    []telemetryObjectBlueZoneCustomOptions `json:"-"`
	BlueZoneCustomOptionsRaw string                                 `json:"blueZoneCustomOptions"`
}

func NewLogMatchStart(raw json.RawMessage) (l *LogMatchStart, err error) {
	l = &LogMatchStart{}
	err = json.Unmarshal(raw, l)
	if err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
		err = json.Unmarshal([]byte(l.BlueZoneCustomOptionsRaw), &l.BlueZoneCustomOptions)
	}
	return
}

type LogObjectDestroy struct {
	telemetryEvent
	Character      telemetryObjectCharacter `json:"character"`
	ObjectType     string                   `json:"objectType"`
	ObjectLocation telemetryObjectLocation  `json:"objectLocation"`
}

func NewLogObjectDestroy(raw json.RawMessage) (l *LogObjectDestroy, err error) {
	l = &LogObjectDestroy{}
	err = json.Unmarshal(raw, l)
	if err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

type LogObjectInteraction struct {
	telemetryEvent
	Character                telemetryObjectCharacter `json:"character"`
	ObjectType               string                   `json:"objectType"`
	ObjectTypeStatus         string                   `json:"objectTypeStatus"`
	ObjectTypeAdditionalInfo []string                 `json:"objectTypeAdditionalInfo"`
	Common                   telemetryObjectCommon    `json:"common"`
}

func NewLogObjectInteraction(raw json.RawMessage) (l *LogObjectInteraction, err error) {
	l = &LogObjectInteraction{}
	err = json.Unmarshal(raw, l)
	if err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

type LogParachuteLanding struct {
	telemetryEvent
	Character telemetryObjectCharacter `json:"character"`
	Distance  float64                  `json:"distance"`
}

func NewLogParachuteLanding(raw json.RawMessage) (l *LogParachuteLanding, err error) {
	l = &LogParachuteLanding{}
	err = json.Unmarshal(raw, l)
	if err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

type LogPhaseChange struct {
	telemetryEvent
	Phase       int     `json:"phase"`
	ElapsedTime float64 `json:"elapsedTime"`
}

func NewLogPhaseChange(raw json.RawMessage) (l *LogPhaseChange, err error) {
	l = &LogPhaseChange{}
	err = json.Unmarshal(raw, l)
	if err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

type LogPlayerAttack struct {
	telemetryEvent
	AttackId             int                      `json:"attackId"`
	FireWeaponStackCount int                      `json:"fireWeaponStackCount"`
	Attacker             telemetryObjectCharacter `json:"attacker"`
	AttackType           string                   `json:"attackType"`
	Weapon               telemetryObjectItem      `json:"weapon"`
	Vehicle              telemetryObjectVehicle   `json:"vehicle"`
}

func NewLogPlayerAttack(raw json.RawMessage) (l *LogPlayerAttack, err error) {
	l = &LogPlayerAttack{}
	err = json.Unmarshal(raw, l)
	if err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

type LogPlayerCreate struct {
	telemetryEvent
	Character telemetryObjectCharacter `json:"character"`
}

func NewLogPlayerCreate(raw json.RawMessage) (l *LogPlayerCreate, err error) {
	l = &LogPlayerCreate{}
	err = json.Unmarshal(raw, l)
	if err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

type LogPlayerDestroyBreachableWall struct {
	telemetryEvent
	Attacker telemetryObjectCharacter `json:"attacker"`
	Weapon   telemetryObjectItem      `json:"weapon"`
}

func NewLogPlayerDestroyBreachableWall(raw json.RawMessage) (l *LogPlayerDestroyBreachableWall, err error) {
	l = &LogPlayerDestroyBreachableWall{}
	err = json.Unmarshal(raw, l)
	if err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

type LogPlayerKill struct {
	telemetryEvent
	AttackId                   int                       `json:"attackId"`
	Killer                     telemetryObjectCharacter  `json:"killer"`
	Victim                     telemetryObjectCharacter  `json:"victim"`
	Assistant                  telemetryObjectCharacter  `json:"assistant"`
	DBNOId                     int                       `json:"dBNOId"`
	DamageReason               string                    `json:"damageReason"`
	DamageTypeCategory         string                    `json:"damageTypeCategory"`
	DamageCauserName           string                    `json:"damageCauserName"`
	DamageCauserAdditionalInfo []string                  `json:"damageCauserAdditionalInfo"`
	VictimWeapon               string                    `json:"VictimWeapon"`
	VictimWeaponAdditionalInfo []string                  `json:"VictimWeaponAdditionalInfo"`
	Distance                   float64                   `json:"distance"`
	VictimGameResult           telemetryObjectGameResult `json:"victimGameResult"`
	IsThroughPenetrableWall    bool                      `json:"isThroughPenetrableWall"`
}

func NewLogPlayerKill(raw json.RawMessage) (l *LogPlayerKill, err error) {
	l = &LogPlayerKill{}
	err = json.Unmarshal(raw, l)
	if err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

type LogPlayerLogin struct {
	telemetryEvent
	AccountId string `json:"accountId"`
}

func NewLogPlayerLogin(raw json.RawMessage) (l *LogPlayerLogin, err error) {
	l = &LogPlayerLogin{}
	err = json.Unmarshal(raw, l)
	if err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

type LogPlayerLogout struct {
	telemetryEvent
	AccountId string `json:"accountId"`
}

func NewLogPlayerLogout(raw json.RawMessage) (l *LogPlayerLogout, err error) {
	l = &LogPlayerLogout{}
	err = json.Unmarshal(raw, l)
	if err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

type LogPlayerMakeGroggy struct {
	telemetryEvent
	AttackId                   int                      `json:"attackId"`
	Attacker                   telemetryObjectCharacter `json:"attacker"`
	Victim                     telemetryObjectCharacter `json:"victim"`
	DamageReason               string                   `json:"damageReason"`
	DamageTypeCategory         string                   `json:"damageTypeCategory"`
	DamageCauserName           string                   `json:"damageCauserName"`
	DamageCauserAdditionalInfo []string                 `json:"damageCauserAdditionalInfo"`
	VictimWeapon               string                   `json:"VictimWeapon"`
	VictimWeaponAdditionalInfo []string                 `json:"VictimWeaponAdditionalInfo"`
	Distance                   float64                  `json:"distance"`
	IsAttackerInVehicle        bool                     `json:"isAttackerInVehicle"`
	DBNOId                     int                      `json:"dBNOId"`
	IsThroughPenetrableWall    bool                     `json:"isThroughPenetrableWall"`
}

func NewLogPlayerMakeGroggy(raw json.RawMessage) (l *LogPlayerMakeGroggy, err error) {
	l = &LogPlayerMakeGroggy{}
	err = json.Unmarshal(raw, l)
	if err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

type LogPlayerPosition struct {
	telemetryEvent
	Character       telemetryObjectCharacter `json:"character"`
	Vehicle         telemetryObjectVehicle   `json:"vehicle"`
	ElapsedTime     float64                  `json:"elapsedTime"`
	NumAlivePlayers int                      `json:"numAlivePlayers"`
}

func NewLogPlayerPosition(raw json.RawMessage) (l *LogPlayerPosition, err error) {
	l = &LogPlayerPosition{}
	err = json.Unmarshal(raw, l)
	if err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

type LogPlayerRevive struct {
	telemetryEvent
	Reviver telemetryObjectCharacter `json:"reviver"`
	Victim  telemetryObjectCharacter `json:"victim"`
	DBNOId  int                      `json:"dBNOId"`
}

func NewLogPlayerRevive(raw json.RawMessage) (l *LogPlayerRevive, err error) {
	l = &LogPlayerRevive{}
	err = json.Unmarshal(raw, l)
	if err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

type LogPlayerTakeDamage struct {
	telemetryEvent
	AttackId           int                      `json:"attackId"`
	Attacker           telemetryObjectCharacter `json:"attacker"`
	Victim             telemetryObjectCharacter `json:"victim"`
	DamageTypeCategory string                   `json:"damageTypeCategory"`
	DamageReason       string                   `json:"damageReason"`
	// 1.0 damage = 1.0 health. Net damage after armor; damage to health
	Damage                  float64 `json:"damage"`
	DamageCauserName        string  `json:"damageCauserName"`
	IsThroughPenetrableWall bool    `json:"isThroughPenetrableWall"`
}

func NewLogPlayerTakeDamage(raw json.RawMessage) (l *LogPlayerTakeDamage, err error) {
	l = &LogPlayerTakeDamage{}
	err = json.Unmarshal(raw, l)
	if err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

type LogPlayerUseThrowable struct {
	telemetryEvent
	AttackId             int                      `json:"attackId"`
	FireWeaponStackCount int                      `json:"fireWeaponStackCount"`
	Attacker             telemetryObjectCharacter `json:"attacker"`
	AttackType           string                   `json:"attackType"`
	Weapon               telemetryObjectItem      `json:"weapon"`
}

func NewLogPlayerUseThrowable(raw json.RawMessage) (l *LogPlayerUseThrowable, err error) {
	l = &LogPlayerUseThrowable{}
	err = json.Unmarshal(raw, l)
	if err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

type LogRedZoneEnded struct {
	telemetryEvent
	Drivers []telemetryObjectCharacter `json:"drivers"`
}

func NewLogRedZoneEnded(raw json.RawMessage) (l *LogRedZoneEnded, err error) {
	l = &LogRedZoneEnded{}
	err = json.Unmarshal(raw, l)
	if err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

type LogSwimEnd struct {
	telemetryEvent
	Character           telemetryObjectCharacter `json:"character"`
	SwimDistance        float64                  `json:"swimDistance"`
	MaxSwimDepthOfWater float64                  `json:"maxSwimDepthOfWater"`
}

func NewLogSwimEnd(raw json.RawMessage) (l *LogSwimEnd, err error) {
	l = &LogSwimEnd{}
	err = json.Unmarshal(raw, l)
	if err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

type LogSwimStart struct {
	telemetryEvent
	Character telemetryObjectCharacter `json:"character"`
}

func NewLogSwimStart(raw json.RawMessage) (l *LogSwimStart, err error) {
	l = &LogSwimStart{}
	err = json.Unmarshal(raw, l)
	if err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

type LogVaultStart struct {
	telemetryEvent
	Character   telemetryObjectCharacter `json:"character"`
	IsLedgeGrab bool                     `json:"isLedgeGrab"`
}

func NewLogVaultStart(raw json.RawMessage) (l *LogVaultStart, err error) {
	l = &LogVaultStart{}
	err = json.Unmarshal(raw, &l)
	if err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

type LogVehicleDestroy struct {
	telemetryEvent
	AttackId           int                      `json:"attackId"`
	Attacker           telemetryObjectCharacter `json:"attacker"`
	Vehicle            telemetryObjectVehicle   `json:"vehicle"`
	DamageTypeCategory string                   `json:"damageTypeCategory"`
	DamageCauserName   string                   `json:"damageCauserName"`
	Distance           float64                  `json:"distance"`
}

func NewLogVehicleDestroy(raw json.RawMessage) (l *LogVehicleDestroy, err error) {
	l = &LogVehicleDestroy{}
	err = json.Unmarshal(raw, l)
	if err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

type LogVehicleLeave struct {
	telemetryEvent
	Character        telemetryObjectCharacter   `json:"character"`
	Vehicle          telemetryObjectVehicle     `json:"vehicle"`
	RideDistance     float64                    `json:"rideDistance"`
	SeatIndex        int                        `json:"seatIndex"`
	MaxSpeed         float64                    `json:"maxSpeed"`
	FellowPassengers []telemetryObjectCharacter `json:"fellowPassengers"`
}

func NewLogVehicleLeave(raw json.RawMessage) (l *LogVehicleLeave, err error) {
	l = &LogVehicleLeave{}
	err = json.Unmarshal(raw, l)
	if err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

type LogVehicleRide struct {
	telemetryEvent
	Character        telemetryObjectCharacter   `json:"character"`
	Vehicle          telemetryObjectVehicle     `json:"vehicle"`
	SeatIndex        int                        `json:"seatIndex"`
	FellowPassengers []telemetryObjectCharacter `json:"fellowPassengers"`
}

func NewLogVehicleRide(raw json.RawMessage) (l *LogVehicleRide, err error) {
	l = &LogVehicleRide{}
	err = json.Unmarshal(raw, l)
	if err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

type LogWeaponFireCount struct {
	telemetryEvent
	Character telemetryObjectCharacter `json:"character"`
	WeaponId  string                   `json:"weaponId"`
	// Increments of 10
	FireCount int `json:"fireCount"`
}

func NewLogWeaponFireCount(raw json.RawMessage) (l *LogWeaponFireCount, err error) {
	l = &LogWeaponFireCount{}
	err = json.Unmarshal(raw, l)
	if err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

type LogWheelDestroy struct {
	telemetryEvent
	AttackId           int                      `json:"attackId"`
	Attacker           telemetryObjectCharacter `json:"attacker"`
	Vehicle            telemetryObjectVehicle   `json:"vehicle"`
	DamageTypeCategory string                   `json:"damageTypeCategory"`
	DamageCauserName   string                   `json:"damageCauserName"`
}

func NewLogWheelDestroy(raw json.RawMessage) (l *LogWheelDestroy, err error) {
	l = &LogWheelDestroy{}
	err = json.Unmarshal(raw, l)
	if err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}
