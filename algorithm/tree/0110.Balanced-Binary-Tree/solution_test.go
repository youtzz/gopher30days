package problem0110

import (
	"github.com/sevenger/gopher30days/algorithm"
	"testing"
)

func Test_isBalanced(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test_isBalanced 1",
			args: args{
				algorithm.NewBinaryTreeByArgs(3, 9, 20, algorithm.NULL, algorithm.NULL, 15, 7),
			},
			want: true,
		},
		{
			name: "Test_isBalanced 2",
			args: args{
				algorithm.NewBinaryTreeByArgs(1, 2, 2, 3, 3, algorithm.NULL, algorithm.NULL, 4, 4),
			},
			want: false,
		},
		{
			name: "Test_isBalanced 3",
			args: args{
				nil,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced(tt.args.root); got != tt.want {
				t.Errorf("isBalanced() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isBalanced_best(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test_isBalanced_best 1",
			args: args{
				algorithm.NewBinaryTreeByArgs(3, 9, 20, algorithm.NULL, algorithm.NULL, 15, 7),
			},
			want: true,
		},
		{
			name: "Test_isBalanced_best 2",
			args: args{
				algorithm.NewBinaryTreeByArgs(1, 2, 2, 3, 3, algorithm.NULL, algorithm.NULL, 4, 4),
			},
			want: false,
		},
		{
			name: "Test_isBalanced_best 3",
			args: args{
				nil,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isBalanced_best(tt.args.root); got != tt.want {
				t.Errorf("isBalanced_best() = %v, want %v", got, tt.want)
			}
		})
	}
}
