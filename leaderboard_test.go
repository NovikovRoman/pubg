package pubg

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

func TestClient_Leaderboards(t *testing.T) {
	c := NewClient(os.Getenv("APIKEY"), nil)

	leaderboards, err := c.Leaderboards(SteamPlatform, testSeasonID, DuoMode, 1)
	require.Nil(t, err)

	require.Equal(t, leaderboards.Data.Attributes.GameMode, DuoMode)
	require.Equal(t, leaderboards.Data.Attributes.ShardId, SteamPlatform)

	require.True(t, len(leaderboards.Data.Relationships.Players.Data) > 0)
	require.True(t, len(leaderboards.Included) > 0)

	// Когда проходим все тесты, необходимо притормаживать, тк ограничение 10 запросов в минуту
	if os.Getenv("PAUSE") != "" {
		time.Sleep(pause)
	}
}
