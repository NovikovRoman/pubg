package pubg

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
)

func TestClient_Matches(t *testing.T) {
	c := NewClient(os.Getenv("APIKEY"), nil)
	matches, err := c.Matches(SteamPlatform, testMatchID)
	require.Nil(t, err)

	require.Equal(t, matches.Data.Attributes.ShardId, SteamPlatform)
	require.Equal(t, matches.Data.ID, testMatchID)
	require.True(t, len(matches.Assets) > 0)
	require.True(t, len(matches.Rosters) > 0)
	require.True(t, len(matches.Participants) > 0)
}
