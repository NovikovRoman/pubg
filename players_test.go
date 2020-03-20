package pubg

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestClient_Player(t *testing.T) {
	player, err := cTest.Player(SteamPlatform, testAccountID)
	require.Nil(t, err)
	require.True(t, len(player.Data.Relationships.Matches.Data) > 0)
	require.True(t, player.Data.Attributes.Name != "")

	time.Sleep(pause)
}

func TestClient_PlayersByNames(t *testing.T) {
	names := []string{testAccountName, testAccountName2}
	players, err := cTest.PlayersByNames(SteamPlatform, names...)
	require.Nil(t, err)
	require.Len(t, players.Data, 2)

	// Может вернуть не по порядку
	for _, d := range players.Data {
		if d.ID == testAccountID {
			require.Equal(t, d.ID, testAccountID)
			require.Equal(t, d.Attributes.Name, testAccountName)
			require.True(t, len(d.Relationships.Matches.Data) > 0)

		} else {
			require.Equal(t, d.ID, testAccountID2)
			require.Equal(t, d.Attributes.Name, testAccountName2)
			require.True(t, len(d.Relationships.Matches.Data) > 0)
		}
	}

	time.Sleep(pause)
}

func TestClient_PlayersByIDs(t *testing.T) {
	players, err := cTest.PlayersByIDs(SteamPlatform, testAccountID, testAccountID2)
	require.Nil(t, err)
	require.Len(t, players.Data, 2)

	// Может вернуть не по порядку
	for _, d := range players.Data {
		if d.ID == testAccountID {
			require.Equal(t, d.ID, testAccountID)
			require.Equal(t, d.Attributes.Name, testAccountName)
			require.True(t, len(d.Relationships.Matches.Data) > 0)

		} else {
			require.Equal(t, d.ID, testAccountID2)
			require.Equal(t, d.Attributes.Name, testAccountName2)
			require.True(t, len(d.Relationships.Matches.Data) > 0)
		}
	}

	time.Sleep(pause)
}
