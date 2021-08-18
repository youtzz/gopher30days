package problem0705

import (
	"testing"
)

func TestMyHashSet(t *testing.T) {
	hashSet := Constructor()

	hashSet.Add(1)
	hashSet.Add(2)

	if got := hashSet.Contains(1); got != true {
		t.Errorf("MyHashSet.Contains(1) = %v, want = %v", got, true)
	}

	if got := hashSet.Contains(3); got != false {
		t.Errorf("MyHashSet.Contains(3) = %v, want = %v", got, false)
	}

	hashSet.Add(2)
	if got := hashSet.Contains(2); got != true {
		t.Errorf("MyHashSet.Contains(2) = %v, want = %v", got, true)
	}

	hashSet.Remove(2)
	if got := hashSet.Contains(2); got != false {
		t.Errorf("MyHashSet.Contains(2) = %v, want = %v", got, false)
	}

}
