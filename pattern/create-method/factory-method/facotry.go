package factory_method

import "fmt"

// Factory 抽象工厂
type Factory interface {
	New() Game
}

// Game 抽象产品
type Game interface {
	Play()
}

// MobaGameFactory 具体工厂
type MobaGameFactory struct{}

func (*MobaGameFactory) New() Game {
	return &LOL{}
}

// LOL 具体产品
type LOL struct{}

func (*LOL) Play() {
	fmt.Println("de ma xi ya wan sui")
}

type FpsGameFactory struct{}

func (*FpsGameFactory) New() Game {
	return &CS{}
}

type CS struct{}

func (*CS) Play() {
	fmt.Println("rush b")
}
