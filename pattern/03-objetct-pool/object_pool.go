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

// Get 从池中获取对象
func (p Pool) Get() *Object {
	return <-p
}

// Put 把对象返还给池
func (p Pool) Put(obj *Object) {
	p <- obj
}
