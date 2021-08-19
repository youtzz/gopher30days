package problem0021

import (
	"github.com/sevenger/gopher30days/algorithm"
	"reflect"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "testcase 1",
			args: args{algorithm.NewLinkedListByArgs(1, 2, 4),
				algorithm.NewLinkedListByArgs(1, 3, 4)},
			want: algorithm.NewLinkedListByArgs(1, 1, 2, 3, 4, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mergeTwoLists_recusive(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "testcase 1",
			args: args{algorithm.NewLinkedListByArgs(1, 2, 4),
				algorithm.NewLinkedListByArgs(1, 3, 4)},
			want: algorithm.NewLinkedListByArgs(1, 1, 2, 3, 4, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeTwoLists_recusive(tt.args.l1, tt.args.l2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeTwoLists_recusive() = %v, want %v", got, tt.want)
			}
		})
	}
}
