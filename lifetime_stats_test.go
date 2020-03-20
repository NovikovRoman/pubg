package pubg

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestClient_LifetimeStatsPlayer(t *testing.T) {
	stats, err := cTest.LifetimeStatsPlayer(SteamPlatform, testAccountID)
	require.Nil(t, err)
	require.True(t, len(stats.Data.Attributes.GameModeStats) > 0)

	require.True(t, len(stats.Data.Relationships.MatchesSoloFPP.Data) >= 0)
	require.Equal(t, stats.Data.Relationships.Player.Data.ID, testAccountID)

	time.Sleep(pause)
}

func TestClient_LifetimeStatsPlayers(t *testing.T) {
	stats, err := cTest.LifetimeStatsPlayers(SteamPlatform, DuoMode, testAccountID, testAccountID2)
	require.Nil(t, err)
	require.Len(t, stats.Data, 2)

	// Может вернуть не по порядку
	for _, d := range stats.Data {
		if d.Relationships.Player.Data.ID != testAccountID {
			require.Equal(t, d.Relationships.Player.Data.ID, testAccountID2)
		}
		require.True(t, len(d.Attributes.GameModeStats) > 0)
	}

	time.Sleep(pause)
}
