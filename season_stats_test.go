package pubg

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestClient_Seasons(t *testing.T) {
	seasons, err := cTest.Seasons(SteamPlatform)
	require.Nil(t, err)
	require.True(t, len(seasons.Data) > 0)

	time.Sleep(pause)
}

func TestClient_RankedStatsPlayer(t *testing.T) {
	stats, err := cTest.RankedStatsPlayer(SteamPlatform, testSeasonID, testAccountID)
	require.Nil(t, err)

	if len(stats.Data.Attributes.RankedGameModeStats) > 0 {
		for _, s := range stats.Data.Attributes.RankedGameModeStats {
			require.True(t, s.BestRankPoint > 0)
		}
	}

	time.Sleep(pause)
}

func TestClient_SeasonStatsPlayer(t *testing.T) {
	stats, err := cTest.SeasonStatsPlayer(SteamPlatform, testSeasonID, testAccountID)
	require.Nil(t, err)
	require.True(t, len(stats.Data.Attributes.GameModeStats) > 0)

	require.True(t, len(stats.Data.Relationships.MatchesSoloFPP.Data) >= 0)
	require.Equal(t, stats.Data.Relationships.Player.Data.ID, testAccountID)
	require.Equal(t, stats.Data.Relationships.Season.Data.ID, testSeasonID)

	time.Sleep(pause)
}

func TestClient_SeasonStatsPlayers(t *testing.T) {
	stats, err := cTest.SeasonStatsPlayers(SteamPlatform, testSeasonID, DuoMode, testAccountID, testAccountID2)
	require.Nil(t, err)
	require.Len(t, stats.Data, 2)

	// Может вернуть не по порядку
	for _, d := range stats.Data {
		if d.Relationships.Player.Data.ID != testAccountID {
			require.Equal(t, d.Relationships.Player.Data.ID, testAccountID2)
		}

		require.True(t, len(d.Attributes.GameModeStats) > 0)
		require.Equal(t, d.Relationships.Season.Data.ID, testSeasonID)
	}

	time.Sleep(pause)
}
