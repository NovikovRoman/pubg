package pubg

type Platform struct {
	code string
}

func (p Platform) String() string {
	return p.code
}

func SteamPlatform() Platform {
	return Platform{code: "steam"}
}

func PsnPlatform() Platform {
	return Platform{code: "psn"}
}

func XboxPlatform() Platform {
	return Platform{code: "xbox"}
}

func KakaoPlatform() Platform {
	return Platform{code: "kakao"}
}
