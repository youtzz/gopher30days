package objetct_pool

import "testing"

func Test(t *testing.T) {
	pool := NewPool(100)
	if len(*pool) != 100 {
		t.Error("error")
	}
	for ob := range *pool {
		ob.Do()
	}
}
