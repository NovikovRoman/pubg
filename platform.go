package pubg

const (
	EmptyPlatform      = ""
	SteamPlatform      = "steam"
	PsnPlatform        = "psn"
	XboxPlatform       = "xbox"
	KakaoPlatform      = "kakao"
	ConsolePlatform    = "console"
	TournamentPlatform = "tournament"
)

type Platform string

func (p Platform) IsEmpty() bool {
	return p == EmptyPlatform
}
