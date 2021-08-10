package _1_builder

type Builder interface {
	BuildHead()
	BuildBody()
	GetResult() string
}
