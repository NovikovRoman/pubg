package pubg

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"testing"
	"time"
)

func TestClient_NewTelemetryFromURL(t *testing.T) {
	tm, err := NewTelemetryFromURL(testTelemetryURL, nil)
	require.Nil(t, err)

	for {
		eventDate, eventType, raw, ok, err := tm.Next()
		require.Nil(t, err)

		if !ok {
			break
		}

		testEvents(t, eventDate, eventType, raw)
	}
}

func TestClient_NewTelemetryFromFile(t *testing.T) {
	tm, err := NewTelemetryFromFile(testTelemetryFile)
	require.Nil(t, err)

	for {
		eventDate, eventType, raw, ok, err := tm.Next()
		require.Nil(t, err)

		if !ok {
			break
		}

		testEvents(t, eventDate, eventType, raw)
	}
}

func TestClient_NewTelemetryFromBytes(t *testing.T) {
	f, err := ioutil.ReadFile(testTelemetryFile)
	require.Nil(t, err)

	tm, err := NewTelemetryFromBytes(f)
	require.Nil(t, err)

	for {
		eventDate, eventType, raw, ok, err := tm.Next()
		require.Nil(t, err)

		if !ok {
			break
		}

		testEvents(t, eventDate, eventType, raw)
	}
}

func testEvents(t *testing.T, eventDate time.Time, eventType string, raw json.RawMessage) {
	switch eventType {
	case "LogArmorDestroy":
		l, err := NewLogArmorDestroy(raw)
		require.Nil(t, err)
		require.Equal(t, l.Date, eventDate)

	case "LogBlackZoneEnded":
		l, err := NewLogBlackZoneEnded(raw)
		require.Nil(t, err)
		require.Equal(t, l.Date, eventDate)

	case "LogCarePackageLand":
		l, err := NewLogCarePackageLand(raw)
		require.Nil(t, err)
		require.Equal(t, l.Date, eventDate)

	case "LogCarePackageSpawn":
		l, err := NewLogCarePackageSpawn(raw)
		require.Nil(t, err)
		require.Equal(t, l.Date, eventDate)

	case "LogGameStatePeriodic":
		l, err := NewLogGameStatePeriodic(raw)
		require.Nil(t, err)
		require.Equal(t, l.Date, eventDate)

	case "LogHeal":
		l, err := NewLogHeal(raw)
		require.Nil(t, err)
		require.Equal(t, l.Date, eventDate)

	case "LogItemAttach":
		l, err := NewLogItemAttach(raw)
		require.Nil(t, err)
		require.Equal(t, l.Date, eventDate)

	case "LogItemDetach":
		l, err := NewLogItemDetach(raw)
		require.Nil(t, err)
		require.Equal(t, l.Date, eventDate)

	case "LogItemDrop":
		l, err := NewLogItemDrop(raw)
		require.Nil(t, err)
		require.Equal(t, l.Date, eventDate)

	case "LogItemEquip":
		l, err := NewLogItemEquip(raw)
		require.Nil(t, err)
		require.Equal(t, l.Date, eventDate)

	case "LogItemPickup":
		l, err := NewLogItemPickup(raw)
		require.Nil(t, err)
		require.Equal(t, l.Date, eventDate)

	case "LogItemPickupFromCarepackage":
		l, err := NewLogItemPickupFromCarepackage(raw)
		require.Nil(t, err)
		require.Equal(t, l.Date, eventDate)

	case "LogItemPickupFromCustomPackage":
		l, err := NewLogItemPickupFromCustomPackage(raw)
		require.Nil(t, err)
		require.Equal(t, l.Date, eventDate)

	case "LogItemPickupFromLootBox":
		l, err := NewLogItemPickupFromLootBox(raw)
		require.Nil(t, err)
		require.Equal(t, l.Date, eventDate)

	case "LogItemUnequip":
		l, err := NewLogItemUnequip(raw)
		require.Nil(t, err)
		require.Equal(t, l.Date, eventDate)

	case "LogItemUse":
		l, err := NewLogItemUse(raw)
		require.Nil(t, err)
		require.Equal(t, l.Date, eventDate)

	case "LogMatchDefinition":
		l, err := NewLogMatchDefinition(raw)
		require.Nil(t, err)
		require.Equal(t, l.Date, eventDate)

	case "LogMatchEnd":
		l, err := NewLogMatchEnd(raw)
		require.Nil(t, err)
		require.Equal(t, l.Date, eventDate)

	case "LogMatchStart":
		l, err := NewLogMatchStart(raw)
		require.Nil(t, err)
		require.Equal(t, l.Date, eventDate)

	case "LogObjectDestroy":
		l, err := NewLogObjectDestroy(raw)
		require.Nil(t, err)
		require.Equal(t, l.Date, eventDate)

	case "LogObjectInteraction":
		l, err := NewLogObjectInteraction(raw)
		require.Nil(t, err)
		require.Equal(t, l.Date, eventDate)

	case "LogParachuteLanding":
		l, err := NewLogParachuteLanding(raw)
		require.Nil(t, err)
		require.Equal(t, l.Date, eventDate)

	case "LogPhaseChange":
		l, err := NewLogPhaseChange(raw)
		require.Nil(t, err)
		require.Equal(t, l.Date, eventDate)

	case "LogPlayerAttack":
		l, err := NewLogPlayerAttack(raw)
		require.Nil(t, err)
		require.Equal(t, l.Date, eventDate)

	case "LogPlayerCreate":
		l, err := NewLogPlayerCreate(raw)
		require.Nil(t, err)
		require.Equal(t, l.Date, eventDate)

	case "LogPlayerDestroyBreachableWall":
		l, err := NewLogPlayerDestroyBreachableWall(raw)
		require.Nil(t, err)
		require.Equal(t, l.Date, eventDate)

	case "LogPlayerKill":
		l, err := NewLogPlayerKill(raw)
		require.Nil(t, err)
		require.Equal(t, l.Date, eventDate)

	case "LogPlayerKillV2":
		l, err := NewLogPlayerKillV2(raw)
		require.Nil(t, err)
		require.Equal(t, l.Date, eventDate)

	case "LogPlayerLogin":
		l, err := NewLogPlayerLogin(raw)
		require.Nil(t, err)
		require.Equal(t, l.Date, eventDate)

	case "LogPlayerLogout":
		l, err := NewLogPlayerLogout(raw)
		require.Nil(t, err)
		require.Equal(t, l.Date, eventDate)

	case "LogPlayerMakeGroggy":
		l, err := NewLogPlayerMakeGroggy(raw)
		require.Nil(t, err)
		require.Equal(t, l.Date, eventDate)

	case "LogPlayerPosition":
		l, err := NewLogPlayerPosition(raw)
		require.Nil(t, err)
		require.Equal(t, l.Date, eventDate)

	case "LogPlayerRevive":
		l, err := NewLogPlayerRevive(raw)
		require.Nil(t, err)
		require.Equal(t, l.Date, eventDate)

	case "LogPlayerTakeDamage":
		l, err := NewLogPlayerTakeDamage(raw)
		require.Nil(t, err)
		require.Equal(t, l.Date, eventDate)

	case "LogPlayerUseFlareGun":
		l, err := NewLogPlayerUseFlareGun(raw)
		require.Nil(t, err)
		require.Equal(t, l.Date, eventDate)

	case "LogPlayerUseThrowable":
		l, err := NewLogPlayerUseThrowable(raw)
		require.Nil(t, err)
		require.Equal(t, l.Date, eventDate)

	case "LogRedZoneEnded":
		l, err := NewLogRedZoneEnded(raw)
		require.Nil(t, err)
		require.Equal(t, l.Date, eventDate)

	case "LogSwimEnd":
		l, err := NewLogSwimEnd(raw)
		require.Nil(t, err)
		require.Equal(t, l.Date, eventDate)

	case "LogSwimStart":
		l, err := NewLogSwimStart(raw)
		require.Nil(t, err)
		require.Equal(t, l.Date, eventDate)

	case "LogVaultStart":
		l, err := NewLogVaultStart(raw)
		require.Nil(t, err)
		require.Equal(t, l.Date, eventDate)

	case "LogVehicleDestroy":
		l, err := NewLogVehicleDestroy(raw)
		require.Nil(t, err)
		require.Equal(t, l.Date, eventDate)

	case "LogVehicleDamage":
		l, err := NewLogVehicleDamage(raw)
		require.Nil(t, err)
		require.Equal(t, l.Date, eventDate)

	case "LogVehicleLeave":
		l, err := NewLogVehicleLeave(raw)
		require.Nil(t, err)
		require.Equal(t, l.Date, eventDate)

	case "LogVehicleRide":
		l, err := NewLogVehicleRide(raw)
		require.Nil(t, err)
		require.Equal(t, l.Date, eventDate)

	case "LogWeaponFireCount":
		l, err := NewLogWeaponFireCount(raw)
		require.Nil(t, err)
		require.Equal(t, l.Date, eventDate)

	case "LogWheelDestroy":
		l, err := NewLogWheelDestroy(raw)
		require.Nil(t, err)
		require.Equal(t, l.Date, eventDate)

	default:
		t.Error("Event " + eventType + " not found.")
	}
}
