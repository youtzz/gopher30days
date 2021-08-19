package problem0206

import (
	"github.com/sevenger/gopher30days/algorithm"
	"reflect"
	"testing"
)

func Test_reverseList(t *testing.T) {
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
				head: algorithm.NewLinkedListByArgs(1, 2, 3, 4, 5),
			},
			want: algorithm.NewLinkedListByArgs(5, 4, 3, 2, 1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_reverseList_recur(t *testing.T) {
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
				head: algorithm.NewLinkedListByArgs(1, 2, 3, 4, 5),
			},
			want: algorithm.NewLinkedListByArgs(5, 4, 3, 2, 1),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseList_recur(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseList() = %v, want %v", got, tt.want)
			}
		})
	}
}
