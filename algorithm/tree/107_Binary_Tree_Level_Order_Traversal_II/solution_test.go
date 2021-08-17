package Binary_Tree_Level_Order_Traversal_II

import (
	"github.com/sevenger/gopher30days/algorithm"
	"reflect"
	"testing"
)

func Test_levelOrder(t *testing.T) {
	type args struct {
		root *algorithm.TreeNode
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "testcase 1",
			args: args{
				root: algorithm.NewBinaryTreeByArgs("3", "9", "20", "null", "null", "15", "7"),
			},
			want: [][]int{
				{15, 7},
				{9, 20},
				{3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := levelOrderBottom(tt.args.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
