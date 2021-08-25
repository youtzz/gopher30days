package problem0101

import (
	"github.com/sevenger/gopher30days/algorithm"
	"testing"
)

const null = algorithm.NULL

func Test_isSymmetric(t *testing.T) {
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
				root: algorithm.NewBinaryTreeByArgs(1, 2, 2, 3, 4, 4, 3),
			},
			want: true,
		},
		{
			name: "testcase 2",
			args: args{
				root: algorithm.NewBinaryTreeByArgs(1, 2, 2, null, 3, null, 3),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetric(tt.args.root); got != tt.want {
				t.Errorf("isSymmetric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isSymmetric_best(t *testing.T) {
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
				root: algorithm.NewBinaryTreeByArgs(1, 2, 2, 3, 4, 4, 3),
			},
			want: true,
		},
		{
			name: "testcase 2",
			args: args{
				root: algorithm.NewBinaryTreeByArgs(1, 2, 2, null, 3, null, 3),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetric_best(tt.args.root); got != tt.want {
				t.Errorf("isSymmetric_best() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isSymmetric_best_recur(t *testing.T) {
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
				root: algorithm.NewBinaryTreeByArgs(1, 2, 2, 3, 4, 4, 3),
			},
			want: true,
		},
		{
			name: "testcase 2",
			args: args{
				root: algorithm.NewBinaryTreeByArgs(1, 2, 2, null, 3, null, 3),
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isSymmetric_best_recur(tt.args.root); got != tt.want {
				t.Errorf("isSymmetric_best_recur() = %v, want %v", got, tt.want)
			}
		})
	}
}
