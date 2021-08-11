package factory_method

import (
	"testing"
)

func TestAbstractFactory(t *testing.T) {
	var factory Factory
	var game Game

	factory = &MobaGameFactory{}
	game = factory.New()
	game.Play()

	factory = &FpsGameFactory{}
	game = factory.New()
	game.Play()
}
