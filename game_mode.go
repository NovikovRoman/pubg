package pubg

type GameMode struct {
	code string
}

func (m GameMode) String() string {
	return m.code
}

func DuoGameMode() GameMode {
	return GameMode{code: "duo"}
}

func DuoFPPGameMode() GameMode {
	return GameMode{code: "duo-fpp"}
}

func SoloGameMode() GameMode {
	return GameMode{code: "solo"}
}

func SoloFPPGameMode() GameMode {
	return GameMode{code: "solo-fpp"}
}

func SquadGameMode() GameMode {
	return GameMode{code: "squad"}
}

func SquadFPPGameMode() GameMode {
	return GameMode{code: "squad-fpp"}
}
