package simgle_factory_mothod

import (
	"testing"
)

func TestGetFactory(t *testing.T) {
	simpleFactory := SimpleFactory{}

	actionGame, _ := simpleFactory.GetGame(ActionGameType)
	actionGame.Start()

	fpsGame, _ := simpleFactory.GetGame(FpsGameType)
	fpsGame.Start()
}
