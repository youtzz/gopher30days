package problem0142

import (
	"github.com/sevenger/gopher30days/algorithm"
	"reflect"
	"testing"
)

func closure(testcase int) (*algorithm.ListNode, *algorithm.ListNode) {
	var head *algorithm.ListNode
	var want *ListNode
	switch testcase {
	case 1:
		head = algorithm.NewCycleLinedList([]int{1, 2, 3, 4}, 1)
		want = head.Next
	case 2:
		head = algorithm.NewCycleLinedList([]int{1}, 0)
		want = head
	case 3:
		head = algorithm.NewCycleLinedList([]int{1, 2, 3}, -1)
		want = nil
	}
	return head, want
}

func Test_detectCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "testcase 1",
			args: args{
				func() *ListNode {
					head, _ := closure(1)
					return head
				}(),
			},
			want: func() *ListNode {
				_, want := closure(1)
				return want
			}(),
		},
		{
			name: "testcase 2",
			args: args{
				func() *ListNode {
					head, _ := closure(2)
					return head
				}(),
			},
			want: func() *ListNode {
				_, want := closure(2)
				return want
			}(),
		},
		{
			name: "testcase 3",
			args: args{
				func() *ListNode {
					head, _ := closure(3)
					return head
				}(),
			},
			want: func() *ListNode {
				_, want := closure(3)
				return want
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectCycle(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("detectCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
