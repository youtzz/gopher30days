package _1_builder

type Builder interface {
	BuildHead(string)
	BuildBody(string)
	GetResult() string
}
