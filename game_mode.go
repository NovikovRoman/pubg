package pubg

import "errors"

const (
	SoloMode     = "solo"
	SoloFPPMode  = "solo-fpp"
	DuoMode      = "duo"
	DuoFPPMode   = "duo-fpp"
	SquadMode    = "squad"
	SquadFPPMode = "squad-fpp"
)

type GameMode string
// IsValid returns true if the mode is valid.
func (m GameMode) IsValid() bool {
	return m == SoloMode || m == SoloFPPMode || m == DuoMode || m == DuoFPPMode || m == SquadMode || m == SquadFPPMode
}
// TransformToGameMode transforms a string into a GameMode structure.
func TransformToGameMode(name string) (gameMode GameMode, err error) {
	gameMode = SoloMode

	switch name {
	case SoloMode:
		return

	case SoloFPPMode:
		gameMode = SoloFPPMode
		return

	case DuoMode:
		gameMode = DuoMode
		return

	case DuoFPPMode:
		gameMode = DuoFPPMode
		return

	case SquadMode:
		gameMode = SquadMode
		return

	case SquadFPPMode:
		gameMode = SquadFPPMode
		return
	}

	err = errors.New("Game mode not found. ")
	return
}
