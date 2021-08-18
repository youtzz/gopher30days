package problem0098

import (
	"github.com/sevenger/gopher30days/algorithm"
	"testing"
)

func Test_isValidBST(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "testcase 1",
			args: args{
				root: algorithm.NewBinaryTreeByArgs("2", "1", "3"),
			},
			want: true,
		},
		{
			name: "testcase 2",
			args: args{
				root: algorithm.NewBinaryTreeByArgs("5", "1", "4", "null", "null", "3", "6"),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST(tt.args.root); got != tt.want {
				t.Errorf("isValidBST() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValidBST_dac(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "testcase 1",
			args: args{
				root: algorithm.NewBinaryTreeByArgs("2", "1", "3"),
			},
			want: true,
		},
		{
			name: "testcase 2",
			args: args{
				root: algorithm.NewBinaryTreeByArgs("5", "1", "4", "null", "null", "3", "6"),
			},
			want: false,
		},
		{
			name: "testcase 3",
			args: args{
				root: algorithm.NewBinaryTreeByArr([]int{2, 2, 2}),
			},
			want: false,
		},
		{
			name: "testcase 4",
			args: args{
				root: algorithm.NewBinaryTreeByArr([]int{0, -1}),
			},
			want: true,
		},
		{
			name: "testcase 5",
			args: args{
				root: algorithm.NewBinaryTreeByArr([]int{1, algorithm.NULL, 1}),
			},
			want: false,
		},
		{
			name: "testcase 6",
			args: args{
				root: algorithm.NewBinaryTreeByArr([]int{5, 4, 6, algorithm.NULL, algorithm.NULL, 3, 7}),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValidBST_dac(tt.args.root); got != tt.want {
				t.Errorf("isValidBST_dac() = %v, want %v", got, tt.want)
			}
		})
	}
}
