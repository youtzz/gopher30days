package Lowest_Common_Ancestor_of_a_Binary_Tree

import (
	"github.com/sevenger/gopher30days/algorithm"
	"reflect"
	"testing"
)

var root = algorithm.NewBinaryTreeByArgs(1, 2)

func Test_lowestCommonAncestor(t *testing.T) {
	type args struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "testcase 1",
			args: args{
				root: root,
				p:    root,
				q:    root.Left,
			},
			want: root,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.args.root, tt.args.p, tt.args.q); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("lowestCommonAncestor() = %v, want %v", got, tt.want)
			}
		})
	}
}
