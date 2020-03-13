package pubg

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

func TestClient_Tournaments(t *testing.T) {
	c := NewClient(os.Getenv("APIKEY"), nil)
	tournaments, err := c.Tournaments()

	require.Nil(t, err)

	require.True(t, len(tournaments.Data) > 0)
	require.True(t, tournaments.Data[0].ID != "")
	require.Equal(t, tournaments.Data[0].Type, "tournament")
	require.False(t, tournaments.Data[0].Attributes.CreatedAt.IsZero())

	// Когда проходим все тесты, необходимо притормаживать, тк ограничение 10 запросов в минуту
	if os.Getenv("pause") != "" {
		time.Sleep(pause)
	}
}

func TestClient_Tournament(t *testing.T) {
	c := NewClient(os.Getenv("APIKEY"), nil)
	tournament, err := c.Tournament(testTournamentID)
	require.Nil(t, err)

	require.True(t, len(tournament.Included) > 0)
	require.True(t, tournament.Data.ID != "")
	require.Equal(t, tournament.Data.Type, "tournament")
	require.False(t, tournament.Included[0].Attributes.CreatedAt.IsZero())

	// Когда проходим все тесты, необходимо притормаживать, тк ограничение 10 запросов в минуту
	if os.Getenv("pause") != "" {
		time.Sleep(pause)
	}
}
