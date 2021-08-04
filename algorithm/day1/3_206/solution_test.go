package main

import (
	. "github.com/sevenger/gopher30days/algorithm"
	"reflect"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	 testcase := []struct {
		Input *ListNode
		Output *ListNode
	} {
	 	{},
	 }

	for _, v := range testcase {
		if !CompareLinkList(deleteDuplicates(v.Input), v.Output) {
			t.Fatalf("%d %d")
		}
	}
}

func Test_reverseList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}