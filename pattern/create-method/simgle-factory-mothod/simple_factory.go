package simgle_factory_mothod

import (
	"errors"
	"fmt"
)

type GameType int

const (
	ActionGameType GameType = 1 << iota
	FpsGameType
)

type Game interface {
	Start()
}

type ActionGame struct{}

func (g *ActionGame) Start() {
	fmt.Println("死或生启动中...")
}

type FpsGame struct{}

func (g *FpsGame) Start() {
	fmt.Println("沙滩排球启动中...")
}

type SimpleFactory struct{}

func (*SimpleFactory) GetGame(gameType GameType) (Game, error) {
	switch gameType {
	case ActionGameType:
		return &ActionGame{}, nil
	case FpsGameType:
		return &FpsGame{}, nil
	default:
		return nil, errors.New("not support game type")
	}
}
