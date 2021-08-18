package problem0701

import (
	"github.com/sevenger/gopher30days/algorithm"
	"reflect"
	"testing"
)

const null = algorithm.NULL

func Test_insertIntoBST(t *testing.T) {
	type args struct {
		root *TreeNode
		val  int
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "testcase 1",
			args: args{
				root: algorithm.NewBinaryTreeByArgs(4, 2, 7, 1, 3),
				val:  5,
			},
			want: algorithm.NewBinaryTreeByArgs(4, 2, 7, 1, 3, 5),
		},
		{
			name: "testcase 2",
			args: args{
				root: algorithm.NewBinaryTreeByArgs(40, 20, 60, 10, 30, 50, 70),
				val:  25,
			},
			want: algorithm.NewBinaryTreeByArgs(40, 20, 60, 10, 30, 50, 70, null, null, 25),
		},
		{
			name: "testcase 3",
			args: args{
				root: algorithm.NewBinaryTreeByArgs(4, 2, 7, 1, 3, null, null, null, null, null, null),
				val:  5,
			},
			want: algorithm.NewBinaryTreeByArgs(4, 2, 7, 1, 3, 5),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertIntoBST(tt.args.root, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertIntoBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
