package pubg

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClient_GameMode(t *testing.T) {
	var m GameMode
	m = SoloMode
	require.Equal(t, string(m), SoloMode)

	m = SoloFPPMode
	require.Equal(t, string(m), SoloFPPMode)

	m = DuoMode
	require.Equal(t, string(m), DuoMode)

	m = DuoFPPMode
	require.Equal(t, string(m), DuoFPPMode)

	m = SquadMode
	require.Equal(t, string(m), SquadMode)

	m = SquadFPPMode
	require.Equal(t, string(m), SquadFPPMode)
}
