package pubg

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

func TestClient_Seasons(t *testing.T) {
	c := NewClient(os.Getenv("APIKEY"))
	seasons, err := c.Seasons(SteamPlatform)
	require.Nil(t, err)
	require.True(t, len(seasons.Data) > 0)

	// Когда проходим все тесты, необходимо притормаживать, тк ограничение 10 запросов в минуту
	if os.Getenv("pause") != "" {
		time.Sleep(pause)
	}
}

func TestClient_SeasonStatsPlayer(t *testing.T) {
	c := NewClient(os.Getenv("APIKEY"))
	stats, err := c.SeasonStatsPlayer(SteamPlatform, testSeasonID, testAccountID)
	require.Nil(t, err)
	require.True(t, len(stats.Data.Attributes.GameModeStats) > 0)

	require.True(t, len(stats.Data.Relationships.MatchesSoloFPP.Data) >= 0)
	require.Equal(t, stats.Data.Relationships.Player.Data.ID, testAccountID)
	require.Equal(t, stats.Data.Relationships.Season.Data.ID, testSeasonID)

	// Когда проходим все тесты, необходимо притормаживать, тк ограничение 10 запросов в минуту
	if os.Getenv("pause") != "" {
		time.Sleep(pause)
	}
}

func TestClient_SeasonStatsPlayers(t *testing.T) {
	c := NewClient(os.Getenv("APIKEY"))
	stats, err := c.SeasonStatsPlayers(SteamPlatform, testSeasonID, DuoMode, testAccountID, testAccountID2)
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

	// Когда проходим все тесты, необходимо притормаживать, тк ограничение 10 запросов в минуту
	if os.Getenv("pause") != "" {
		time.Sleep(pause)
	}
}
