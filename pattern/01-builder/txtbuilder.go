package _1_builder

type TxtBuilder struct {
	content string
}

func (t TxtBuilder) BuildHead(head string) {
	t.content = head
}

func (t TxtBuilder) BuildBody(body string) {
	t.content += body
}

func (t TxtBuilder) GetResult() string {
	return t.content
}
