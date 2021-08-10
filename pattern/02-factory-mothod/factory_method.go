package factory_mothod

import "errors"

type GameType int

const (
	ActionGame GameType = 1 << iota
	FpsGame
	RpgGame
)

type Game struct {
	Name string
}

func GetGame(gameType GameType) *Game {
	var game *Game
	switch gameType {
	case ActionGame:
		game = NewActionGame()
	case FpsGame:
		game = NewFpsGame()
	case RpgGame:
		game = NewRpgGame()
	default:
		panic(errors.New("unknown type"))
	}
	return game
}

func NewActionGame() *Game {
	return &Game{Name: "死或生"}
}

func NewFpsGame() *Game {
	return &Game{Name: "沙滩排球"}
}

func NewRpgGame() *Game {
	return &Game{Name: "莱莎炼金工坊"}
}
