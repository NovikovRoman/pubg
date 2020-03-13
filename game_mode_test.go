package pubg

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClient_GameMode(t *testing.T) {
	var (
		m   GameMode
		err error
	)

	m = SoloMode
	require.True(t, m.IsValid())
	require.Equal(t, string(m), SoloMode)
	m, err = TransformToGameMode(SoloMode)
	require.Nil(t, err)
	require.Equal(t, string(m), SoloMode)

	m = SoloFPPMode
	require.True(t, m.IsValid())
	require.Equal(t, string(m), SoloFPPMode)
	m, err = TransformToGameMode(SoloFPPMode)
	require.Nil(t, err)
	require.Equal(t, string(m), SoloFPPMode)

	m = DuoMode
	require.True(t, m.IsValid())
	require.Equal(t, string(m), DuoMode)
	m, err = TransformToGameMode(DuoMode)
	require.Nil(t, err)
	require.Equal(t, string(m), DuoMode)

	m = DuoFPPMode
	require.True(t, m.IsValid())
	require.Equal(t, string(m), DuoFPPMode)
	m, err = TransformToGameMode(DuoFPPMode)
	require.Nil(t, err)
	require.Equal(t, string(m), DuoFPPMode)

	m = SquadMode
	require.True(t, m.IsValid())
	require.Equal(t, string(m), SquadMode)
	m, err = TransformToGameMode(SquadMode)
	require.Nil(t, err)
	require.Equal(t, string(m), SquadMode)

	m = SquadFPPMode
	require.True(t, m.IsValid())
	require.Equal(t, string(m), SquadFPPMode)
	m, err = TransformToGameMode(SquadFPPMode)
	require.Nil(t, err)
	require.Equal(t, string(m), SquadFPPMode)

	m = "unknown"
	require.False(t, m.IsValid())
	require.Equal(t, string(m), "unknown")
	m, err = TransformToGameMode(string(m))
	require.NotNil(t, err)
	require.Equal(t, string(m), SoloMode)
}
