package pubg

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClient_Platform(t *testing.T) {
	var p Platform
	p = EmptyPlatform
	require.True(t, p.IsEmpty())
	require.Equal(t, string(p), EmptyPlatform)

	p = SteamPlatform
	require.False(t, p.IsEmpty())
	require.Equal(t, string(p), "steam")

	p = PsnPlatform
	require.False(t, p.IsEmpty())
	require.Equal(t, string(p), "psn")

	p = XboxPlatform
	require.False(t, p.IsEmpty())
	require.Equal(t, string(p), "xbox")

	p = KakaoPlatform
	require.False(t, p.IsEmpty())
	require.Equal(t, string(p), "kakao")

	p = ConsolePlatform
	require.False(t, p.IsEmpty())
	require.Equal(t, string(p), "console")

	p = TournamentPlatform
	require.False(t, p.IsEmpty())
	require.Equal(t, string(p), "tournament")
}
