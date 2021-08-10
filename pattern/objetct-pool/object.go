package objetct_pool

import "fmt"

type Object struct {
	Name string
}

func (o *Object) Do() {
	fmt.Println(o.Name)
}

type Pool chan *Object

func NewPool(count int) *Pool {
	pool := make(Pool, count)
	defer close(pool)
	for i := 0; i < count; i++ {
		pool <- &Object{}
	}
	return &pool
}
