package singleton

import (
	"testing"
)

func TestGetInstance(t *testing.T) {
	for i := 0; i < 100; i++ {
		go GetInstance()
	}
}
