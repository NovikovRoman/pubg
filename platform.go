package pubg

type Platform struct {
	code string
}

func (p Platform) String() string {
	return p.code
}

// Steam
func SteamPlatform() Platform {
	return Platform{code: "steam"}
}

// PS4
func PsnPlatform() Platform {
	return Platform{code: "psn"}
}

// Xbox
func XboxPlatform() Platform {
	return Platform{code: "xbox"}
}

// Kakao
func KakaoPlatform() Platform {
	return Platform{code: "kakao"}
}

// PS4/Xbox (used for the /matches and /samples endpoints)
func ConsolePlatform() Platform {
	return Platform{code: "console"}
}

// Tournaments
func TournamentPlatform() Platform {
	return Platform{code: "tournament"}
}
