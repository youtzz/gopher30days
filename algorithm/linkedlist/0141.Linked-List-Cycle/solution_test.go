package problem0141

import (
	"github.com/sevenger/gopher30days/algorithm"
	"testing"
)

func Test_hasCycle(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "testcase 1",
			args: args{
				head: algorithm.NewCycleLinedList([]int{1, 2, 3, 4}, 1),
			},
			want: true,
		},
		{
			name: "testcase 2",
			args: args{
				head: algorithm.NewCycleLinedList([]int{1, 2, 3, 4}, -1),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hasCycle(tt.args.head); got != tt.want {
				t.Errorf("hasCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
