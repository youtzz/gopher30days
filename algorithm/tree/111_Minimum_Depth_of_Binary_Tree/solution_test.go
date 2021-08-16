package Minimum_Depth_of_Binary_Tree

import (
	. "github.com/sevenger/gopher30days/algorithm"
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
				NewBinaryTreeByArgs("3", "9", "20", "null", "null", "15", "7"),
			},
			want: 2,
		}, {
			name: "Test_minDepth 2",
			args: args{
				NewBinaryTreeByArgs("2", "null", "3", "null", "4", "null", "5", "null", "6"),
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
				NewBinaryTreeByArgs("3", "9", "20", "null", "null", "15", "7"),
			},
			want: 2,
		}, {
			name: "Test_minDepth 2",
			args: args{
				NewBinaryTreeByArgs("2", "null", "3", "null", "4", "null", "5", "null", "6"),
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
