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

// LogArmorDestroy structure.
type LogArmorDestroy struct {
	telemetryEvent
	AttackID           int                      `json:"attackId"`
	Attacker           telemetryObjectCharacter `json:"attacker"`
	Victim             telemetryObjectCharacter `json:"victim"`
	DamageTypeCategory string                   `json:"damageTypeCategory"`
	DamageReason       string                   `json:"damageReason"`
	DamageCauserName   string                   `json:"damageCauserName"`
	Item               telemetryObjectItem      `json:"item"`
	Distance           float64                  `json:"distance"`
}

// NewLogArmorDestroy create new LogArmorDestroy structure.
func NewLogArmorDestroy(raw json.RawMessage) (l *LogArmorDestroy, err error) {
	l = &LogArmorDestroy{}
	if err = json.Unmarshal(raw, l); err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

// LogBlackZoneEnded structure.
type LogBlackZoneEnded struct {
	telemetryEvent
	Survivors []telemetryObjectCharacter `json:"survivors"`
}

// NewLogBlackZoneEnded create new LogBlackZoneEnded structure.
func NewLogBlackZoneEnded(raw json.RawMessage) (l *LogBlackZoneEnded, err error) {
	l = &LogBlackZoneEnded{}
	if err = json.Unmarshal(raw, l); err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

// LogCarePackageLand structure.
type LogCarePackageLand struct {
	telemetryEvent
	ItemPackage ItemPackage `json:"itemPackage"`
}

// NewLogCarePackageLand create new LogCarePackageLand structure.
func NewLogCarePackageLand(raw json.RawMessage) (l *LogCarePackageLand, err error) {
	l = &LogCarePackageLand{}
	if err = json.Unmarshal(raw, l); err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

// LogCarePackageSpawn structure.
type LogCarePackageSpawn struct {
	telemetryEvent
	ItemPackage ItemPackage `json:"itemPackage"`
}

// NewLogCarePackageSpawn create new LogCarePackageSpawn structure.
func NewLogCarePackageSpawn(raw json.RawMessage) (l *LogCarePackageSpawn, err error) {
	l = &LogCarePackageSpawn{}
	if err = json.Unmarshal(raw, l); err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

// LogGameStatePeriodic structure.
type LogGameStatePeriodic struct {
	telemetryEvent
	GameState telemetryObjectGameState `json:"gameState"`
}

// NewLogGameStatePeriodic create new LogGameStatePeriodic structure.
func NewLogGameStatePeriodic(raw json.RawMessage) (l *LogGameStatePeriodic, err error) {
	l = &LogGameStatePeriodic{}
	if err = json.Unmarshal(raw, l); err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

// LogHeal structure.
type LogHeal struct {
	telemetryEvent
	Character  telemetryObjectCharacter `json:"character"`
	Item       telemetryObjectItem      `json:"item"`
	HealAmount float64                  `json:"healAmount"`
}

// NewLogHeal create new LogHeal structure.
func NewLogHeal(raw json.RawMessage) (l *LogHeal, err error) {
	l = &LogHeal{}
	if err = json.Unmarshal(raw, l); err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

// LogItemAttach structure.
type LogItemAttach struct {
	telemetryEvent
	Character  telemetryObjectCharacter `json:"character"`
	ParentItem telemetryObjectItem      `json:"parentItem"`
	ChildItem  telemetryObjectItem      `json:"childItem"`
}

// NewLogItemAttach create new LogItemAttach structure.
func NewLogItemAttach(raw json.RawMessage) (l *LogItemAttach, err error) {
	l = &LogItemAttach{}
	if err = json.Unmarshal(raw, l); err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

// LogItemDetach structure.
type LogItemDetach struct {
	telemetryEvent
	Character  telemetryObjectCharacter `json:"character"`
	ParentItem telemetryObjectItem      `json:"parentItem"`
	ChildItem  telemetryObjectItem      `json:"childItem"`
}

// NewLogItemDetach create new LogItemDetach structure.
func NewLogItemDetach(raw json.RawMessage) (l *LogItemDetach, err error) {
	l = &LogItemDetach{}
	if err = json.Unmarshal(raw, l); err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

// LogItemDrop structure.
type LogItemDrop struct {
	telemetryEvent
	Character telemetryObjectCharacter `json:"character"`
	Item      telemetryObjectItem      `json:"item"`
}

// NewLogItemDrop create new LogItemDrop structure.
func NewLogItemDrop(raw json.RawMessage) (l *LogItemDrop, err error) {
	l = &LogItemDrop{}
	if err = json.Unmarshal(raw, l); err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

// LogItemEquip structure.
type LogItemEquip struct {
	telemetryEvent
	Character telemetryObjectCharacter `json:"character"`
	Item      telemetryObjectItem      `json:"item"`
}

// NewLogItemEquip create new LogItemEquip structure.
func NewLogItemEquip(raw json.RawMessage) (l *LogItemEquip, err error) {
	l = &LogItemEquip{}
	if err = json.Unmarshal(raw, l); err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

// LogItemPickup structure.
type LogItemPickup struct {
	telemetryEvent
	Character telemetryObjectCharacter `json:"character"`
	Item      telemetryObjectItem      `json:"item"`
}

// NewLogItemPickup create new LogItemPickup structure.
func NewLogItemPickup(raw json.RawMessage) (l *LogItemPickup, err error) {
	l = &LogItemPickup{}
	if err = json.Unmarshal(raw, l); err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

// LogItemPickupFromCarepackage structure.
type LogItemPickupFromCarepackage struct {
	telemetryEvent
	Character           telemetryObjectCharacter `json:"character"`
	Item                telemetryObjectItem      `json:"item"`
	CarePackageUniqueId float64                  `json:"carePackageUniqueId"`
}

// NewLogItemPickupFromCarepackage create new LogItemPickupFromCarepackage structure.
func NewLogItemPickupFromCarepackage(raw json.RawMessage) (l *LogItemPickupFromCarepackage, err error) {
	l = &LogItemPickupFromCarepackage{}
	if err = json.Unmarshal(raw, l); err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

// LogItemPickupFromCustomPackage structure.
type LogItemPickupFromCustomPackage struct {
	telemetryEvent
	Character telemetryObjectCharacter `json:"character"`
	Item      telemetryObjectItem      `json:"item"`
}

// NewLogItemPickupFromCustomPackage create new LogItemPickupFromCustomPackage structure.
func NewLogItemPickupFromCustomPackage(raw json.RawMessage) (l *LogItemPickupFromCustomPackage, err error) {
	l = &LogItemPickupFromCustomPackage{}
	if err = json.Unmarshal(raw, l); err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

// LogItemPickupFromLootBox structure.
type LogItemPickupFromLootBox struct {
	telemetryEvent
	Character        telemetryObjectCharacter `json:"character"`
	Item             telemetryObjectItem      `json:"item"`
	OwnerTeamId      int                      `json:"ownerTeamId"`
	CreatorAccountId string                   `json:"creatorAccountId"`
}

// NewLogItemPickupFromLootBox create new LogItemPickupFromLootBox structure.
func NewLogItemPickupFromLootBox(raw json.RawMessage) (l *LogItemPickupFromLootBox, err error) {
	l = &LogItemPickupFromLootBox{}
	if err = json.Unmarshal(raw, l); err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

// LogItemUnequip structure.
type LogItemUnequip struct {
	telemetryEvent
	Character telemetryObjectCharacter `json:"character"`
	Item      telemetryObjectItem      `json:"item"`
}

// NewLogItemUnequip create new LogItemUnequip structure.
func NewLogItemUnequip(raw json.RawMessage) (l *LogItemUnequip, err error) {
	l = &LogItemUnequip{}
	if err = json.Unmarshal(raw, l); err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

// LogItemUse structure.
type LogItemUse struct {
	telemetryEvent
	Character telemetryObjectCharacter `json:"character"`
	Item      telemetryObjectItem      `json:"item"`
}

// NewLogItemUse create new LogItemUse structure.
func NewLogItemUse(raw json.RawMessage) (l *LogItemUse, err error) {
	l = &LogItemUse{}
	if err = json.Unmarshal(raw, l); err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

// LogMatchDefinition structure.
type LogMatchDefinition struct {
	telemetryEvent
	MatchId     string `json:"MatchId"`
	SeasonState string `json:"SeasonState"`
}

// NewLogMatchDefinition create new LogMatchDefinition structure.
func NewLogMatchDefinition(raw json.RawMessage) (l *LogMatchDefinition, err error) {
	l = &LogMatchDefinition{}
	if err = json.Unmarshal(raw, l); err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

// LogMatchEnd structure.
type LogMatchEnd struct {
	telemetryEvent
	Characters []telemetryObjectCharacterWrapper `json:"characters"`
	// Shows winning players only
	GameResultOnFinished telemetryObjectGameResultOnFinished `json:"gameResultOnFinished"`
}

// NewLogMatchEnd create new LogMatchEnd structure.
func NewLogMatchEnd(raw json.RawMessage) (l *LogMatchEnd, err error) {
	l = &LogMatchEnd{}
	if err = json.Unmarshal(raw, l); err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

// LogMatchStart structure.
type LogMatchStart struct {
	telemetryEvent
	MapName                  string                                 `json:"mapName"`
	WeatherId                string                                 `json:"weatherId"`
	Characters               []telemetryObjectCharacterWrapper      `json:"characters"`
	CameraViewBehaviour      string                                 `json:"cameraViewBehaviour"`
	TeamSize                 int                                    `json:"teamSize"`
	IsCustomGame             bool                                   `json:"isCustomGame"`
	IsEventMode              bool                                   `json:"isEventMode"`
	BlueZoneCustomOptions    []telemetryObjectBlueZoneCustomOptions `json:"-"`
	BlueZoneCustomOptionsRaw string                                 `json:"blueZoneCustomOptions"`
}

// NewLogMatchStart create new LogMatchStart structure.
func NewLogMatchStart(raw json.RawMessage) (l *LogMatchStart, err error) {
	l = &LogMatchStart{}
	if err = json.Unmarshal(raw, l); err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
		err = json.Unmarshal([]byte(l.BlueZoneCustomOptionsRaw), &l.BlueZoneCustomOptions)
	}
	return
}

// LogObjectDestroy structure.
type LogObjectDestroy struct {
	telemetryEvent
	Character      telemetryObjectCharacter `json:"character"`
	ObjectType     string                   `json:"objectType"`
	ObjectLocation telemetryObjectLocation  `json:"objectLocation"`
}

// NewLogObjectDestroy create new LogObjectDestroy structure.
func NewLogObjectDestroy(raw json.RawMessage) (l *LogObjectDestroy, err error) {
	l = &LogObjectDestroy{}
	if err = json.Unmarshal(raw, l); err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

// LogObjectInteraction structure.
type LogObjectInteraction struct {
	telemetryEvent
	Character                telemetryObjectCharacter `json:"character"`
	ObjectType               string                   `json:"objectType"`
	ObjectTypeStatus         string                   `json:"objectTypeStatus"`
	ObjectTypeAdditionalInfo []map[string]interface{} `json:"objectTypeAdditionalInfo"`
	Common                   telemetryObjectCommon    `json:"common"`
}

// NewLogObjectInteraction create new LogObjectInteraction structure.
func NewLogObjectInteraction(raw json.RawMessage) (l *LogObjectInteraction, err error) {
	l = &LogObjectInteraction{}
	if err = json.Unmarshal(raw, l); err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

// LogParachuteLanding structure.
type LogParachuteLanding struct {
	telemetryEvent
	Character telemetryObjectCharacter `json:"character"`
	Distance  float64                  `json:"distance"`
}

// NewLogParachuteLanding create new LogParachuteLanding structure.
func NewLogParachuteLanding(raw json.RawMessage) (l *LogParachuteLanding, err error) {
	l = &LogParachuteLanding{}
	if err = json.Unmarshal(raw, l); err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

// LogPhaseChange structure.
type LogPhaseChange struct {
	telemetryEvent
	Phase       int     `json:"phase"`
	ElapsedTime float64 `json:"elapsedTime"`
}

// NewLogPhaseChange create new LogPhaseChange structure.
func NewLogPhaseChange(raw json.RawMessage) (l *LogPhaseChange, err error) {
	l = &LogPhaseChange{}
	if err = json.Unmarshal(raw, l); err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

// LogPlayerAttack structure.
type LogPlayerAttack struct {
	telemetryEvent
	AttackId             int                      `json:"attackId"`
	FireWeaponStackCount int                      `json:"fireWeaponStackCount"`
	Attacker             telemetryObjectCharacter `json:"attacker"`
	AttackType           string                   `json:"attackType"`
	Weapon               telemetryObjectItem      `json:"weapon"`
	Vehicle              telemetryObjectVehicle   `json:"vehicle"`
}

// NewLogPlayerAttack create new LogPlayerAttack structure.
func NewLogPlayerAttack(raw json.RawMessage) (l *LogPlayerAttack, err error) {
	l = &LogPlayerAttack{}
	if err = json.Unmarshal(raw, l); err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

// LogPlayerCreate structure.
type LogPlayerCreate struct {
	telemetryEvent
	Character telemetryObjectCharacter `json:"character"`
}

// NewLogPlayerCreate create new LogPlayerCreate structure.
func NewLogPlayerCreate(raw json.RawMessage) (l *LogPlayerCreate, err error) {
	l = &LogPlayerCreate{}
	if err = json.Unmarshal(raw, l); err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

// LogPlayerDestroyBreachableWall structure.
type LogPlayerDestroyBreachableWall struct {
	telemetryEvent
	Attacker telemetryObjectCharacter `json:"attacker"`
	Weapon   telemetryObjectItem      `json:"weapon"`
}

// NewLogPlayerDestroyBreachableWall create new LogPlayerDestroyBreachableWall structure.
func NewLogPlayerDestroyBreachableWall(raw json.RawMessage) (l *LogPlayerDestroyBreachableWall, err error) {
	l = &LogPlayerDestroyBreachableWall{}
	if err = json.Unmarshal(raw, l); err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

// LogPlayerKill structure.
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

// NewLogPlayerKill create new LogPlayerKill structure.
func NewLogPlayerKill(raw json.RawMessage) (l *LogPlayerKill, err error) {
	l = &LogPlayerKill{}
	if err = json.Unmarshal(raw, l); err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

// LogPlayerKillV2 structure.
type LogPlayerKillV2 struct {
	telemetryEvent
	AttackId                   int                       `json:"attackId"`
	DBNOId                     int                       `json:"dBNOId"`
	VictimGameResult           telemetryObjectGameResult `json:"victimGameResult"`
	Victim                     telemetryObjectCharacter  `json:"victim"`
	VictimWeapon               string                    `json:"VictimWeapon"`
	VictimWeaponAdditionalInfo []string                  `json:"VictimWeaponAdditionalInfo"`
	DBNOMaker                  telemetryObjectCharacter  `json:"dBNOMaker"`
	DBNODamageInfo             telemetryObjectDamageInfo `json:"dBNODamageInfo"`
	Finisher                   telemetryObjectCharacter  `json:"finisher"`
	FinishDamageInfo           telemetryObjectDamageInfo `json:"finishDamageInfo"`
	Killer                     telemetryObjectCharacter  `json:"killer"`
	KillerDamageInfo           telemetryObjectDamageInfo `json:"killerDamageInfo"`
	IsSuicide                  bool                      `json:"isSuicide"`
}

// NewLogPlayerKillV2 create new LogPlayerKillV2 structure.
func NewLogPlayerKillV2(raw json.RawMessage) (l *LogPlayerKillV2, err error) {
	l = &LogPlayerKillV2{}
	if err = json.Unmarshal(raw, l); err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

// LogPlayerLogin structure.
type LogPlayerLogin struct {
	telemetryEvent
	AccountId string `json:"accountId"`
}

// NewLogPlayerLogin create new LogPlayerLogin structure.
func NewLogPlayerLogin(raw json.RawMessage) (l *LogPlayerLogin, err error) {
	l = &LogPlayerLogin{}
	if err = json.Unmarshal(raw, l); err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

// LogPlayerLogout structure.
type LogPlayerLogout struct {
	telemetryEvent
	AccountId string `json:"accountId"`
}

// NewLogPlayerLogout create new LogPlayerLogout structure.
func NewLogPlayerLogout(raw json.RawMessage) (l *LogPlayerLogout, err error) {
	l = &LogPlayerLogout{}
	if err = json.Unmarshal(raw, l); err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

// LogPlayerMakeGroggy structure.
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

// NewLogPlayerMakeGroggy create new LogPlayerMakeGroggy structure.
func NewLogPlayerMakeGroggy(raw json.RawMessage) (l *LogPlayerMakeGroggy, err error) {
	l = &LogPlayerMakeGroggy{}
	if err = json.Unmarshal(raw, l); err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

// LogPlayerPosition structure.
type LogPlayerPosition struct {
	telemetryEvent
	Character       telemetryObjectCharacter `json:"character"`
	Vehicle         telemetryObjectVehicle   `json:"vehicle"`
	ElapsedTime     float64                  `json:"elapsedTime"`
	NumAlivePlayers int                      `json:"numAlivePlayers"`
}

// NewLogPlayerPosition create new LogPlayerPosition structure.
func NewLogPlayerPosition(raw json.RawMessage) (l *LogPlayerPosition, err error) {
	l = &LogPlayerPosition{}
	if err = json.Unmarshal(raw, l); err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

// LogPlayerRevive structure.
type LogPlayerRevive struct {
	telemetryEvent
	Reviver telemetryObjectCharacter `json:"reviver"`
	Victim  telemetryObjectCharacter `json:"victim"`
	DBNOId  int                      `json:"dBNOId"`
}

// NewLogPlayerRevive create new LogPlayerRevive structure.
func NewLogPlayerRevive(raw json.RawMessage) (l *LogPlayerRevive, err error) {
	l = &LogPlayerRevive{}
	if err = json.Unmarshal(raw, l); err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

// LogPlayerTakeDamage structure.
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

// NewLogPlayerTakeDamage create new LogPlayerTakeDamage structure.
func NewLogPlayerTakeDamage(raw json.RawMessage) (l *LogPlayerTakeDamage, err error) {
	l = &LogPlayerTakeDamage{}
	if err = json.Unmarshal(raw, l); err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

// LogPlayerUseFlareGun structure.
type LogPlayerUseFlareGun struct {
	telemetryEvent
	AttackId             int                      `json:"attackId"`
	FireWeaponStackCount int                      `json:"fireWeaponStackCount"`
	Attacker             telemetryObjectCharacter `json:"attacker"`
	AttackType           string                   `json:"attackType"`
	Weapon               telemetryObjectItem      `json:"weapon"`
}

// NewLogPlayerUseFlareGun create new LogPlayerUseFlareGun structure.
func NewLogPlayerUseFlareGun(raw json.RawMessage) (l *LogPlayerUseFlareGun, err error) {
	l = &LogPlayerUseFlareGun{}
	if err = json.Unmarshal(raw, l); err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

// LogPlayerUseThrowable structure.
type LogPlayerUseThrowable struct {
	telemetryEvent
	AttackId             int                      `json:"attackId"`
	FireWeaponStackCount int                      `json:"fireWeaponStackCount"`
	Attacker             telemetryObjectCharacter `json:"attacker"`
	AttackType           string                   `json:"attackType"`
	Weapon               telemetryObjectItem      `json:"weapon"`
}

// NewLogPlayerUseThrowable create new LogPlayerUseThrowable structure.
func NewLogPlayerUseThrowable(raw json.RawMessage) (l *LogPlayerUseThrowable, err error) {
	l = &LogPlayerUseThrowable{}
	if err = json.Unmarshal(raw, l); err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

// LogRedZoneEnded structure.
type LogRedZoneEnded struct {
	telemetryEvent
	Drivers []telemetryObjectCharacter `json:"drivers"`
}

// NewLogRedZoneEnded create new LogRedZoneEnded structure.
func NewLogRedZoneEnded(raw json.RawMessage) (l *LogRedZoneEnded, err error) {
	l = &LogRedZoneEnded{}
	if err = json.Unmarshal(raw, l); err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

// LogSwimEnd structure.
type LogSwimEnd struct {
	telemetryEvent
	Character           telemetryObjectCharacter `json:"character"`
	SwimDistance        float64                  `json:"swimDistance"`
	MaxSwimDepthOfWater float64                  `json:"maxSwimDepthOfWater"`
}

// NewLogSwimEnd create new LogSwimEnd structure.
func NewLogSwimEnd(raw json.RawMessage) (l *LogSwimEnd, err error) {
	l = &LogSwimEnd{}
	if err = json.Unmarshal(raw, l); err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

// LogSwimStart structure.
type LogSwimStart struct {
	telemetryEvent
	Character telemetryObjectCharacter `json:"character"`
}

// NewLogSwimStart create new LogSwimStart structure.
func NewLogSwimStart(raw json.RawMessage) (l *LogSwimStart, err error) {
	l = &LogSwimStart{}
	if err = json.Unmarshal(raw, l); err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

// LogVaultStart structure.
type LogVaultStart struct {
	telemetryEvent
	Character   telemetryObjectCharacter `json:"character"`
	IsLedgeGrab bool                     `json:"isLedgeGrab"`
}

// NewLogVaultStart create new LogVaultStart structure.
func NewLogVaultStart(raw json.RawMessage) (l *LogVaultStart, err error) {
	l = &LogVaultStart{}
	err = json.Unmarshal(raw, &l)
	if err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

// LogVehicleDestroy structure.
type LogVehicleDestroy struct {
	telemetryEvent
	AttackId           int                      `json:"attackId"`
	Attacker           telemetryObjectCharacter `json:"attacker"`
	Vehicle            telemetryObjectVehicle   `json:"vehicle"`
	DamageTypeCategory string                   `json:"damageTypeCategory"`
	DamageCauserName   string                   `json:"damageCauserName"`
	Distance           float64                  `json:"distance"`
}

// NewLogVehicleDestroy create new LogVehicleDestroy structure.
func NewLogVehicleDestroy(raw json.RawMessage) (l *LogVehicleDestroy, err error) {
	l = &LogVehicleDestroy{}
	if err = json.Unmarshal(raw, l); err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

// LogVehicleDamage structure.
type LogVehicleDamage struct {
	telemetryEvent
	AttackId           int                      `json:"attackId"`
	Attacker           telemetryObjectCharacter `json:"attacker"`
	Vehicle            telemetryObjectVehicle   `json:"vehicle"`
	DamageTypeCategory string                   `json:"damageTypeCategory"`
	DamageCauserName   string                   `json:"damageCauserName"`
	Damage             float64                  `json:"damage"`
	Distance           float64                  `json:"distance"`
}

// NewLogVehicleDamage create new LogVehicleDamage structure.
func NewLogVehicleDamage(raw json.RawMessage) (l *LogVehicleDamage, err error) {
	l = &LogVehicleDamage{}
	if err = json.Unmarshal(raw, l); err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

// LogVehicleLeave structure.
type LogVehicleLeave struct {
	telemetryEvent
	Character        telemetryObjectCharacter   `json:"character"`
	Vehicle          telemetryObjectVehicle     `json:"vehicle"`
	RideDistance     float64                    `json:"rideDistance"`
	SeatIndex        int                        `json:"seatIndex"`
	MaxSpeed         float64                    `json:"maxSpeed"`
	FellowPassengers []telemetryObjectCharacter `json:"fellowPassengers"`
}

// NewLogVehicleLeave create new LogVehicleLeave structure.
func NewLogVehicleLeave(raw json.RawMessage) (l *LogVehicleLeave, err error) {
	l = &LogVehicleLeave{}
	if err = json.Unmarshal(raw, l); err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

// LogVehicleRide structure.
type LogVehicleRide struct {
	telemetryEvent
	Character        telemetryObjectCharacter   `json:"character"`
	Vehicle          telemetryObjectVehicle     `json:"vehicle"`
	SeatIndex        int                        `json:"seatIndex"`
	FellowPassengers []telemetryObjectCharacter `json:"fellowPassengers"`
}

// NewLogVehicleRide create new LogVehicleRide structure.
func NewLogVehicleRide(raw json.RawMessage) (l *LogVehicleRide, err error) {
	l = &LogVehicleRide{}
	if err = json.Unmarshal(raw, l); err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

// LogWeaponFireCount structure.
type LogWeaponFireCount struct {
	telemetryEvent
	Character telemetryObjectCharacter `json:"character"`
	WeaponId  string                   `json:"weaponId"`
	// Increments of 10
	FireCount int `json:"fireCount"`
}

// NewLogWeaponFireCount create new LogWeaponFireCount structure.
func NewLogWeaponFireCount(raw json.RawMessage) (l *LogWeaponFireCount, err error) {
	l = &LogWeaponFireCount{}
	if err = json.Unmarshal(raw, l); err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}

// LogWheelDestroy structure.
type LogWheelDestroy struct {
	telemetryEvent
	AttackId           int                      `json:"attackId"`
	Attacker           telemetryObjectCharacter `json:"attacker"`
	Vehicle            telemetryObjectVehicle   `json:"vehicle"`
	DamageTypeCategory string                   `json:"damageTypeCategory"`
	DamageCauserName   string                   `json:"damageCauserName"`
}

// NewLogWheelDestroy create new LogWheelDestroy structure.
func NewLogWheelDestroy(raw json.RawMessage) (l *LogWheelDestroy, err error) {
	l = &LogWheelDestroy{}
	if err = json.Unmarshal(raw, l); err == nil {
		l.Date, _ = time.Parse(layoutDateTime, l.DateRaw)
	}
	return
}
