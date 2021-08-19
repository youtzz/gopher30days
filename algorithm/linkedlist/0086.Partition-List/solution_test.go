package problem0086

import (
	"github.com/sevenger/gopher30days/algorithm"
	"reflect"
	"testing"
)

func Test_partition(t *testing.T) {
	type args struct {
		head *algorithm.ListNode
		x    int
	}
	tests := []struct {
		name string
		args args
		want *algorithm.ListNode
	}{
		{
			name: "testcase 1",
			args: args{
				head: algorithm.NewLinkedListByArgs(1, 4, 3, 2, 5, 2),
				x:    3,
			},
			want: algorithm.NewLinkedListByArgs(1, 2, 2, 4, 3, 5),
		},
		{
			name: "testcase 2",
			args: args{
				head: algorithm.NewLinkedListByArgs(2, 1),
				x:    2,
			},
			want: algorithm.NewLinkedListByArgs(1, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := partition(tt.args.head, tt.args.x); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("partition() = %v, want %v", got, tt.want)
			}
		})
	}
}
