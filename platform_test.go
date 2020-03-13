package pubg

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestClient_Platform(t *testing.T) {
	var (
		p   Platform
		err error
	)

	p = EmptyPlatform
	require.True(t, p.IsEmpty())
	require.True(t, p.IsValid())
	require.Equal(t, string(p), EmptyPlatform)
	p, err = TransformToPlatform(EmptyPlatform)
	require.Nil(t, err)
	require.Equal(t, string(p), EmptyPlatform)

	p = SteamPlatform
	require.False(t, p.IsEmpty())
	require.True(t, p.IsValid())
	require.Equal(t, string(p), SteamPlatform)
	p, err = TransformToPlatform(SteamPlatform)
	require.Nil(t, err)
	require.Equal(t, string(p), SteamPlatform)

	p = PsnPlatform
	require.False(t, p.IsEmpty())
	require.True(t, p.IsValid())
	require.Equal(t, string(p), PsnPlatform)
	p, err = TransformToPlatform(PsnPlatform)
	require.Nil(t, err)
	require.Equal(t, string(p), PsnPlatform)

	p = XboxPlatform
	require.False(t, p.IsEmpty())
	require.True(t, p.IsValid())
	require.Equal(t, string(p), XboxPlatform)
	p, err = TransformToPlatform(XboxPlatform)
	require.Nil(t, err)
	require.Equal(t, string(p), XboxPlatform)

	p = KakaoPlatform
	require.False(t, p.IsEmpty())
	require.True(t, p.IsValid())
	require.Equal(t, string(p), KakaoPlatform)
	p, err = TransformToPlatform(KakaoPlatform)
	require.Nil(t, err)
	require.Equal(t, string(p), KakaoPlatform)

	p = ConsolePlatform
	require.False(t, p.IsEmpty())
	require.True(t, p.IsValid())
	require.Equal(t, string(p), ConsolePlatform)
	p, err = TransformToPlatform(ConsolePlatform)
	require.Nil(t, err)
	require.Equal(t, string(p), ConsolePlatform)

	p = TournamentPlatform
	require.False(t, p.IsEmpty())
	require.True(t, p.IsValid())
	require.Equal(t, string(p), TournamentPlatform)
	p, err = TransformToPlatform(TournamentPlatform)
	require.Nil(t, err)
	require.Equal(t, string(p), TournamentPlatform)

	p = "unknown"
	require.False(t, p.IsEmpty())
	require.False(t, p.IsValid())
	require.Equal(t, string(p), "unknown")
	p, err = TransformToPlatform(string(p))
	require.NotNil(t, err)
	require.Equal(t, string(p), EmptyPlatform)
}
