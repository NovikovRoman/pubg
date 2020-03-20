package pubg

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestClient_Tournaments(t *testing.T) {
	tournaments, err := cTest.Tournaments()

	require.Nil(t, err)

	require.True(t, len(tournaments.Data) > 0)
	require.True(t, tournaments.Data[0].ID != "")
	require.Equal(t, tournaments.Data[0].Type, "tournament")
	require.False(t, tournaments.Data[0].Attributes.CreatedAt.IsZero())

	time.Sleep(pause)
}

func TestClient_Tournament(t *testing.T) {
	tournament, err := cTest.Tournament(testTournamentID)
	require.Nil(t, err)

	require.True(t, len(tournament.Included) > 0)
	require.True(t, tournament.Data.ID != "")
	require.Equal(t, tournament.Data.Type, "tournament")
	require.False(t, tournament.Included[0].Attributes.CreatedAt.IsZero())

	time.Sleep(pause)
}
