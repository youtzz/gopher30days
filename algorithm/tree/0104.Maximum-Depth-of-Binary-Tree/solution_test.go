package problem0104

import (
	"github.com/sevenger/gopher30days/algorithm"
	"testing"
)

func Test_maxDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testcase 1",
			args: args{
				algorithm.NewBinaryTreeByArgs(3, 9, 20, algorithm.NULL, algorithm.NULL, 15, 7),
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth(tt.args.root); got != tt.want {
				t.Errorf("maxDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxDepth_bfs(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testcase 1",
			args: args{
				algorithm.NewBinaryTreeByArgs(3, 9, 20, algorithm.NULL, algorithm.NULL, 15, 7),
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxDepth_bfs(tt.args.root); got != tt.want {
				t.Errorf("maxDepth_bfs() = %v, want %v", got, tt.want)
			}
		})
	}
}
