package pubg

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestClient_Leaderboards(t *testing.T) {
	leaderboards, err := cTest.Leaderboards(SteamPlatform, testSeasonID, DuoMode, 1)
	if err != nil {
		if err, ok := err.(*ErrBadRequest); ok {
			require.Equal(t, err.GetDetail(), "missing data - ShardID: "+SteamPlatform)
			time.Sleep(pause)
			return
		}
	}
	require.Nil(t, err)

	require.Equal(t, leaderboards.Data.Attributes.GameMode, DuoMode)
	require.Equal(t, leaderboards.Data.Attributes.ShardId, SteamPlatform)

	require.True(t, len(leaderboards.Data.Relationships.Players.Data) > 0)
	require.True(t, len(leaderboards.Included) > 0)

	time.Sleep(pause)
}
