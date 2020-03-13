package pubg

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

func TestClient_LifetimeStatsPlayer(t *testing.T) {
	c := NewClient(os.Getenv("APIKEY"), nil)
	stats, err := c.LifetimeStatsPlayer(SteamPlatform, testAccountID)
	require.Nil(t, err)
	require.True(t, len(stats.Data.Attributes.GameModeStats) > 0)

	require.True(t, len(stats.Data.Relationships.MatchesSoloFPP.Data) >= 0)
	require.Equal(t, stats.Data.Relationships.Player.Data.ID, testAccountID)

	// Когда проходим все тесты, необходимо притормаживать, тк ограничение 10 запросов в минуту
	if os.Getenv("PAUSE") != "" {
		time.Sleep(pause)
	}
}

func TestClient_LifetimeStatsPlayers(t *testing.T) {
	c := NewClient(os.Getenv("APIKEY"), nil)
	stats, err := c.LifetimeStatsPlayers(SteamPlatform, DuoMode, testAccountID, testAccountID2)
	require.Nil(t, err)
	require.Len(t, stats.Data, 2)

	// Может вернуть не по порядку
	for _, d := range stats.Data {
		if d.Relationships.Player.Data.ID != testAccountID {
			require.Equal(t, d.Relationships.Player.Data.ID, testAccountID2)
		}
		require.True(t, len(d.Attributes.GameModeStats) > 0)
	}

	// Когда проходим все тесты, необходимо притормаживать, тк ограничение 10 запросов в минуту
	if os.Getenv("PAUSE") != "" {
		time.Sleep(pause)
	}
}
