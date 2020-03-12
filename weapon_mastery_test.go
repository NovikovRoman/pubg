package pubg

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

func TestClient_WeaponMastery(t *testing.T) {
	c := NewClient(os.Getenv("APIKEY"))
	pSteam := SteamPlatform()
	weaponMastery, err := c.WeaponMastery(pSteam, testAccountID)
	require.Nil(t, err)

	require.Equal(t, weaponMastery.Data.Attributes.Platform, SteamPlatform().String())
	require.Equal(t, weaponMastery.Data.ID, testAccountID)
	require.True(t, len(weaponMastery.Data.Attributes.WeaponSummaries) > 0)

	// Когда проходим все тесты, необходимо притормаживать, тк ограничение 10 запросов в минуту
	if os.Getenv("pause") != "" {
		time.Sleep(pause)
	}
}
