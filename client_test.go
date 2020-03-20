package pubg

import (
	"github.com/stretchr/testify/require"
	"net/http"
	"net/url"
	"os"
	"testing"
)

const (
	testAccountID     = "account.c0e530e9b7244b358def282782f893af"
	testAccountName   = "WackyJacky101"
	testAccountID2    = "account.2f6161f9becd4f8d9bcac25c5f049be8"
	testAccountName2  = "Lecuma46"
	testSeasonID      = "division.bro.official.pc-2018-06"
	testMatchID       = "9c8ca51a-ee05-4ca0-8415-d59f2605639b"
	testTournamentID  = "eu-pgs20"
	testTelemetryFile = "testdata/telemetry.json"
	testTelemetryURL  = "https://telemetry-cdn.playbattlegrounds.com/bluehole-pubg/steam/2020/03/10/21/50/1de3f581-6319-11ea-96a5-ae06d0909089-telemetry.json"
)

func TestNewClient(t *testing.T) {
	if os.Getenv("PROXY") == "" {
		return
	}

	proxyURL, err := url.Parse(os.Getenv("PROXY"))
	require.Nil(t, err)

	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyURL),
	}
	c := NewClient(os.Getenv("APIKEY"), transport)
	require.True(t, c.Status())

	m, err := c.Matches(SteamPlatform, testMatchID)
	require.Nil(t, err)
	require.Equal(t, m.Data.ID, testMatchID)
}
