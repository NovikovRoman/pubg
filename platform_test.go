package pubg

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClient_Platform(t *testing.T) {
	p := EmptyPlatform()
	require.True(t, p.IsEmpty())
	require.Equal(t, p.String(), "")

	p = SteamPlatform()
	require.False(t, p.IsEmpty())
	require.Equal(t, p.String(), "steam")

	p = PsnPlatform()
	require.False(t, p.IsEmpty())
	require.Equal(t, p.String(), "psn")

	p = XboxPlatform()
	require.False(t, p.IsEmpty())
	require.Equal(t, p.String(), "xbox")

	p = KakaoPlatform()
	require.False(t, p.IsEmpty())
	require.Equal(t, p.String(), "kakao")

	p = ConsolePlatform()
	require.False(t, p.IsEmpty())
	require.Equal(t, p.String(), "console")

	p = TournamentPlatform()
	require.False(t, p.IsEmpty())
	require.Equal(t, p.String(), "tournament")
}
