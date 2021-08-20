package cmd

import (
	"fmt"
	"testing"
)

func TestNewSolution(t *testing.T) {
	name := NewSolution("tree", 1, "Insert into binary tree")
	fmt.Println(name)
}
