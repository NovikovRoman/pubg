package pubg

import "errors"

const (
	// EmptyPlatform the platform is not specified
	EmptyPlatform = ""
	// SteamPlatform - Steam
	SteamPlatform = "steam"
	// PsnPlatform - PS4
	PsnPlatform = "psn"
	// XboxPlatform - Xbox
	XboxPlatform = "xbox"
	// KakaoPlatform - Kakao
	KakaoPlatform = "kakao"
	// ConsolePlatform - Console
	ConsolePlatform = "console"
	// TournamentPlatform - Tournament
	TournamentPlatform = "tournament"
	// StadiaPlatform - Stadia
	StadiaPlatform = "stadia"
)

// Platform as string
type Platform string

// IsEmpty returns true if the platform is empty.
func (p Platform) IsEmpty() bool {
	return p == EmptyPlatform
}

// IsValid returns true if the platform is valid.
func (p Platform) IsValid() bool {
	return p == EmptyPlatform || p == SteamPlatform || p == PsnPlatform || p == XboxPlatform ||
		p == KakaoPlatform || p == ConsolePlatform || p == TournamentPlatform || p == StadiaPlatform
}

// TransformToPlatform transforms a string into a Platform structure.
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

	case StadiaPlatform:
		platform = StadiaPlatform
		return
	}

	err = errors.New("Platform not found. ")
	return
}
