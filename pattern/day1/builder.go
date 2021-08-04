package main

import (
	"fmt"
	"strings"
)

type Product struct {
	Name string
	Age  int
	call func()
}

type Builder struct {
	Name string
	Age  int
	call func()
}

func (b *Builder) SetName(name string) *Builder {
	b.Name = name
	return b
}

func (b *Builder) SetAge(age int) *Builder {
	b.Age = age
	return b
}

func (b *Builder) SetCall(call func()) *Builder {
	b.call = call
	return b
}

func (b *Builder) Build() *Product {
	p := Product{
		Name: b.Name,
		Age:  b.Age,
		call: b.call,
	}
	return &p
}

func main() {
	product := (&Builder{}).
		SetName("").
		SetAge(1).
		SetCall(func() {
			fmt.Println("", "2")
		}).
		Build()
	product.call()

	sb := strings.Builder{}
	sb.WriteString("123")
}
