package _3_objetct_pool

import (
	"testing"
)

func TestObjectPool(t *testing.T) {
	pool := NewPool(100)
	if len(pool) != 100 {
		t.Error("Create pool error")
	}

	obj := pool.Get()
	if len(pool) != 99 {
		t.Error("Get obj from pool error")
	}

	obj.DoWork()

	pool.Put(obj)
	if len(pool) != 100 {
		t.Error("Put obj to pool error")
	}
}
