package pubg

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

func TestClient_Player(t *testing.T) {
	c := NewClient(os.Getenv("APIKEY"))
	pSteam := SteamPlatform()
	player, err := c.Player(pSteam, testAccountID)
	require.Nil(t, err)
	require.True(t, len(player.Data.Relationships.Matches.Data) > 0)
	require.True(t, player.Data.Attributes.Name != "")

	// Когда проходим все тесты, необходимо притормаживать, тк ограничение 10 запросов в минуту
	if os.Getenv("pause") != "" {
		time.Sleep(pause)
	}
}

func TestClient_PlayersByNames(t *testing.T) {
	c := NewClient(os.Getenv("APIKEY"))
	pSteam := SteamPlatform()
	names := []string{testAccountName, testAccountName2}
	players, err := c.PlayersByNames(pSteam, names...)
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

	// Когда проходим все тесты, необходимо притормаживать, тк ограничение 10 запросов в минуту
	if os.Getenv("pause") != "" {
		time.Sleep(pause)
	}
}

func TestClient_PlayersByIDs(t *testing.T) {
	c := NewClient(os.Getenv("APIKEY"))
	pSteam := SteamPlatform()
	players, err := c.PlayersByIDs(pSteam, testAccountID, testAccountID2)
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

	// Когда проходим все тесты, необходимо притормаживать, тк ограничение 10 запросов в минуту
	if os.Getenv("pause") != "" {
		time.Sleep(pause)
	}
}
