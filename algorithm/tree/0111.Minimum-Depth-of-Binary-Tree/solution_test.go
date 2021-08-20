package problem0111

import (
	"github.com/sevenger/gopher30days/algorithm"
	"testing"
)

func Test_minDepth(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test_minDepth 1",
			args: args{
				algorithm.NewBinaryTreeByArgs(3, 9, 20, algorithm.NULL, algorithm.NULL, 15, 7),
			},
			want: 2,
		}, {
			name: "Test_minDepth 2",
			args: args{
				algorithm.NewBinaryTreeByArgs(2, algorithm.NULL, 3, algorithm.NULL, 4, algorithm.NULL, 5, algorithm.NULL, 6),
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDepth(tt.args.root); got != tt.want {
				t.Errorf("minDepth() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minDepth_DFS(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test_minDepth 1",
			args: args{
				algorithm.NewBinaryTreeByArgs(3, 9, 20, algorithm.NULL, algorithm.NULL, 15, 7),
			},
			want: 2,
		}, {
			name: "Test_minDepth 2",
			args: args{
				algorithm.NewBinaryTreeByArgs(2, algorithm.NULL, 3, algorithm.NULL, 4, algorithm.NULL, 5, algorithm.NULL, 6),
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minDepth_DFS(tt.args.root); got != tt.want {
				t.Errorf("minDepth_DFS() = %v, want %v", got, tt.want)
			}
		})
	}
}
