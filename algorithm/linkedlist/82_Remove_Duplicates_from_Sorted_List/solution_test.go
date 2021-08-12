package removeDuplicatesFromSortedList

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
		if !CompareLinkedList(deleteDuplicates(v.Input), v.Output) {
			t.Fatalf("%dbfs %dbfs")
		}
	}
}
