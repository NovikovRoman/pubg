package pubg

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestClient_Samples(t *testing.T) {
	samples, err := cTest.Samples(SteamPlatform, time.Now().Add(-time.Hour*48))
	require.Nil(t, err)

	require.Equal(t, samples.Data.Attributes.ShardID, SteamPlatform)
	require.Equal(t, samples.Data.Type, "sample")
	require.True(t, len(samples.Data.Relationships.Matches.Data) > 0)

	time.Sleep(pause)
}
