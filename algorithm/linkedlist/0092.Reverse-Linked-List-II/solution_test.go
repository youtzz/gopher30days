package problem0092

import (
	"github.com/sevenger/gopher30days/algorithm"
	"reflect"
	"testing"
)

func Test_reverseBetween(t *testing.T) {
	type args struct {
		head  *ListNode
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "testcase 1",
			args: args{
				head:  algorithm.NewLinkedListByArgs(1, 2, 3, 4, 5),
				left:  2,
				right: 4,
			},
			want: algorithm.NewLinkedListByArgs(1, 4, 3, 2, 5),
		},
		{
			name: "testcase 2",
			args: args{
				head:  algorithm.NewLinkedListByArgs(5),
				left:  1,
				right: 1,
			},
			want: algorithm.NewLinkedListByArgs(5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBetween(tt.args.head, tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseBetween() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseBetweenBest(t *testing.T) {
	type args struct {
		head  *ListNode
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "testcase 1",
			args: args{
				head:  algorithm.NewLinkedListByArgs(1, 2, 3, 4, 5),
				left:  2,
				right: 4,
			},
			want: algorithm.NewLinkedListByArgs(1, 4, 3, 2, 5),
		},
		{
			name: "testcase 2",
			args: args{
				head:  algorithm.NewLinkedListByArgs(5),
				left:  1,
				right: 1,
			},
			want: algorithm.NewLinkedListByArgs(5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseBetweenBest(tt.args.head, tt.args.left, tt.args.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseBetweenBest() = %v, want %v", got, tt.want)
			}
		})
	}
}
