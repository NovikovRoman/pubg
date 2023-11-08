package pubg

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestClient_SurvivalMastery(t *testing.T) {
	survivalMastery, err := cTest.SurvivalMastery(SteamPlatform, testAccountID)
	require.Nil(t, err)
	require.Equal(t, survivalMastery.Data.ID, testAccountID)
	require.Greater(t, survivalMastery.Data.Attributes.XP, 100)

	time.Sleep(pause)
}
