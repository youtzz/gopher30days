package problem0025

import (
	"github.com/sevenger/gopher30days/algorithm"
	"reflect"
	"testing"
)

func Test_reverseKGroup(t *testing.T) {
	type args struct {
		head *ListNode
		k    int
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
				k:    2,
			},
			want: algorithm.NewLinkedListByArgs(2, 1, 4, 3, 5),
		},
		{
			name: "testcase 2",
			args: args{
				head: algorithm.NewLinkedListByArgs(1, 2, 3, 4, 5),
				k:    3,
			},
			want: algorithm.NewLinkedListByArgs(3, 2, 1, 4, 5),
		},
		{
			name: "testcase 3",
			args: args{
				head: algorithm.NewLinkedListByArgs(1, 2, 3, 4, 5),
				k:    1,
			},
			want: algorithm.NewLinkedListByArgs(1, 2, 3, 4, 5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseKGroup(tt.args.head, tt.args.k); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverseKGroup() = %v, want %v", got, tt.want)
			}
		})
	}
}
