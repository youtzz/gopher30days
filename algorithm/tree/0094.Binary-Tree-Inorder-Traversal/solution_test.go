package problem94

import (
	"github.com/sevenger/gopher30days/algorithm"
	"reflect"
	"testing"
)

type args struct {
	root *TreeNode
}

var tests = []struct {
	name string
	args args
	want []int
}{
	{
		name: "testcase 1",
		args: args{
			root: algorithm.NewBinaryTreeByArgs(1, algorithm.NULL, 2, 3),
		},
		want: []int{1, 3, 2},
	},
	{
		name: "testcase 2",
		args: args{
			root: algorithm.NewBinaryTreeByArgs(),
		},
		want: nil,
	},
	{
		name: "testcase 3",
		args: args{
			root: algorithm.NewBinaryTreeByArgs(1),
		},
		want: []int{1},
	},
	{
		name: "testcase 4",
		args: args{
			root: algorithm.NewBinaryTreeByArgs(1, 2),
		},
		want: []int{2, 1},
	},
	{
		name: "testcase 5",
		args: args{
			root: algorithm.NewBinaryTreeByArgs(1, algorithm.NULL, 2),
		},
		want: []int{1, 2},
	},
}

func Test_inorderTraversal(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotRes := inorderTraversal(tt.args.root); !reflect.DeepEqual(gotRes, tt.want) {
				t.Errorf("inorderTraversal() = %v, want %v", gotRes, tt.want)
			}
		})
	}
}

func Test_inorderTraversal_interactive(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := inorderTraversal_interactive(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("inorderTraversal_interactive() = %v, want %v", got, tt.want)
			}
		})
	}
}
