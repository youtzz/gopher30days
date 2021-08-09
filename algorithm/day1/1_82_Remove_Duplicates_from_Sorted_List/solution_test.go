package main

import (
	. "github.com/sevenger/gopher30days/algorithm"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	testcase := []struct {
		Input  *ListNode
		Output *ListNode
	}{
		{},
	}

	for _, v := range testcase {
		if !CompareLinkList(deleteDuplicates(v.Input), v.Output) {
			t.Fatalf("%d %d")
		}
	}
}
