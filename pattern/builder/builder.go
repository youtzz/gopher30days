package builder

type Builder interface {
	BuildHead(string)
	BuildBody(string)
	GetResult() string
}
