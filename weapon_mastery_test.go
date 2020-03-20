package pubg

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestClient_WeaponMastery(t *testing.T) {
	weaponMastery, err := cTest.WeaponMastery(SteamPlatform, testAccountID)
	require.Nil(t, err)

	require.Equal(t, weaponMastery.Data.Attributes.Platform, SteamPlatform)
	require.Equal(t, weaponMastery.Data.ID, testAccountID)
	require.True(t, len(weaponMastery.Data.Attributes.WeaponSummaries) > 0)

	time.Sleep(pause)
}
