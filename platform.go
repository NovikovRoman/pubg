package pubg

import "errors"

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

func (p Platform) IsValid() bool {
	return p == EmptyPlatform || p == SteamPlatform || p == PsnPlatform || p == XboxPlatform ||
		p == KakaoPlatform || p == ConsolePlatform || p == TournamentPlatform
}

func TransformToPlatform(name string) (platform Platform, err error) {
	platform = EmptyPlatform

	switch name {
	case EmptyPlatform:
		return

	case SteamPlatform:
		platform = SteamPlatform
		return

	case PsnPlatform:
		platform = PsnPlatform
		return

	case XboxPlatform:
		platform = XboxPlatform
		return

	case KakaoPlatform:
		platform = KakaoPlatform
		return

	case ConsolePlatform:
		platform = ConsolePlatform
		return

	case TournamentPlatform:
		platform = TournamentPlatform
		return
	}

	err = errors.New("Platform not found. ")
	return
}
