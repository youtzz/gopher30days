package _3_objetct_pool

import "fmt"

type Object struct {
	Name string
}

func (o *Object) DoWork() {
	fmt.Println(o.Name)
}

type Pool chan *Object

func NewPool(count int) Pool {
	pool := make(Pool, count)
	for i := 0; i < count; i++ {
		pool <- &Object{}
	}
	return pool
}

func (p Pool) Get() *Object {
	return <-p
}

func (p Pool) Put(obj *Object) {
	p <- obj
}
