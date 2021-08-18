package problem0124

import (
	"github.com/sevenger/gopher30days/algorithm"
	"testing"
)

func Test_maxPathSum(t *testing.T) {
	type args struct {
		root *algorithm.TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testcase 1",
			args: args{
				algorithm.NewBinaryTreeByArgs("1", "2", "3"),
			},
			want: 6,
		},
		{
			name: "testcase 2",
			args: args{
				algorithm.NewBinaryTreeByArgs("-10", "9", "20", "null", "null", "15", "7"),
			},
			want: 42,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxPathSum(tt.args.root); got != tt.want {
				t.Errorf("maxPathSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
