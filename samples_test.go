package pubg

import (
	"github.com/stretchr/testify/require"
	"os"
	"testing"
	"time"
)

func TestClient_Samples(t *testing.T) {
	c := NewClient(os.Getenv("APIKEY"))
	samples, err := c.Samples(SteamPlatform, time.Now().Add(-time.Hour*48))
	require.Nil(t, err)

	require.Equal(t, samples.Data.Attributes.ShardId, SteamPlatform)
	require.Equal(t, samples.Data.Type, "sample")
	require.True(t, len(samples.Data.Relationships.Matches.Data) > 0)

	// Когда проходим все тесты, необходимо притормаживать, тк ограничение 10 запросов в минуту
	if os.Getenv("pause") != "" {
		time.Sleep(pause)
	}
}
