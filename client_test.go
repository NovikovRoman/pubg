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
	testTournamentID  = "eu-pgs20"
	testTelemetryFile = "testdata/telemetry.json"
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
