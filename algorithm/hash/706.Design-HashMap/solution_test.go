package problem0706

import (
	"testing"
)

func TestMyHashSet(t *testing.T) {
	hashMap := Constructor()

	hashMap.Put(1, 1)
	hashMap.Put(2, 2)

	if got := hashMap.Get(1); got != 1 {
		t.Errorf("MyHashSet.Get(1) = %v, want = %v", got, 1)
	}

	if got := hashMap.Get(3); got != -1 {
		t.Errorf("MyHashSet.Get(3) = %v, want = %v", got, -1)
	}

	hashMap.Put(2, 1)
	if got := hashMap.Get(2); got != 1 {
		t.Errorf("MyHashSet.Get(2) = %v, want = %v", got, 1)
	}

	hashMap.Remove(2)
	if got := hashMap.Get(2); got != -1 {
		t.Errorf("MyHashSet.Get(2) = %v, want = %v", got, -1)
	}

}
