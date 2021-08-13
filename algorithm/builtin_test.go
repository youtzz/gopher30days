package algorithm

import (
	"reflect"
	"testing"
)

func Test_NewLinkList(t *testing.T) {
	type args struct {
		arg []int
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "build normal link list",
			args: args{arg: []int{1, 2, 3}},
			want: &ListNode{1, &ListNode{2, &ListNode{3, nil}}},
		},
		{
			name: "build nil link list",
			args: args{arg: []int{}},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLinkedList(tt.args.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_CompareLinkList(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"equal", args{l1: NewLinkedList([]int{1, 2, 3}), l2: NewLinkedList([]int{1, 2, 3})}, true},
		{"equal", args{l1: nil, l2: nil}, true},
		{"not equal", args{l1: NewLinkedList([]int{1, 2, 3, 4}), l2: NewLinkedList([]int{1, 2, 3})}, false},
		{"not equal", args{l1: NewLinkedList([]int{1, 2, 3, 4}), l2: nil}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareLinkedList(tt.args.l1, tt.args.l2); got != tt.want {
				t.Errorf("CompareLinkedList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_PrintLinkList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
	}{
		{"1", args{head: NewLinkedList([]int{1, 2, 3, 4, 5})}},
		{"2", args{head: NewLinkedList(nil)}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintLinkedList(tt.args.head)
		})
	}
}

func Test_getFormatLinkListString(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"1", args{head: NewLinkedList([]int{1, 2, 3, 4})}, "{ 1 -> 2 -> 3 -> 4 -> nil }"},
		{"2", args{head: NewLinkedList(nil)}, "{ nil }"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getFormatLinkedListString(tt.args.head); got != tt.want {
				t.Errorf("getFormatLinkedListString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_NewMatrix(t *testing.T) {
	type args struct {
		args [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{
				args: [][]int{
					{1, 3, 5, 7},
					{10, 11, 16, 20},
					{23, 30, 34, 60},
				},
			},
			want: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMatrix(tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewMatrixByRowLen(t *testing.T) {
	type args struct {
		rowLen int
		args   []int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{
				rowLen: 4,
				args: []int{1, 3, 5, 7,
					10, 11, 16, 20,
					23, 30, 34, 60},
			},
			want: [][]int{
				{1, 3, 5, 7},
				{10, 11, 16, 20},
				{23, 30, 34, 60},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMatrixByRowLen(tt.args.rowLen, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMatrixByRowLen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBinaryTree(t *testing.T) {
	type args struct {
		source []string
	}
	tests := []struct {
		name string
		args args
		want *TreeNode
	}{
		{
			name: "testcase 1",
			args: args{
				source: []string{"1", "2", "3"},
			},
			want: &TreeNode{
				Val:   1,
				Left:  &TreeNode{Val: 2},
				Right: &TreeNode{Val: 3},
			},
		},
		{
			name: "testcase 2",
			args: args{
				source: nil,
			},
			want: nil,
		},
		{
			name: "testcase 3",
			args: args{
				source: []string{"1", "nil", "2", "3"},
			},
			want: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val:  2,
					Left: &TreeNode{Val: 3},
				},
			},
		},
		{
			name: "testcase 4",
			args: args{
				source: []string{"1", "nil", "2", "3", "nil", "nil", "4", "5"},
			},
			want: &TreeNode{
				Val: 1,
				Right: &TreeNode{
					Val: 2,
					Left: &TreeNode{
						Val: 3,
						Right: &TreeNode{
							Val:  4,
							Left: &TreeNode{Val: 5},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBinaryTree(tt.args.source); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCompareBinaryTree(t *testing.T) {
	type args struct {
		root1 *TreeNode
		root2 *TreeNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "testcase 1",
			args: args{
				root1: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 2,
					},
				},
				root2: &TreeNode{
					Val: 0,
					Left: &TreeNode{
						Val: 1,
					},
					Right: &TreeNode{
						Val: 2,
					},
				},
			},
			want: true,
		},
		{
			name: "testcase 2",
			args: args{
				root1: &TreeNode{
					Val:   0,
					Left:  nil,
					Right: nil,
				},
				root2: nil,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareBinaryTree(tt.args.root1, tt.args.root2); got != tt.want {
				t.Errorf("CompareBinaryTree() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrintTree(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "testcase 1",
			args: args{
				root: NewBinaryTreeByArgs("1", "nil", "2", "3", "4", "5", "6"),
			},
		},
		{
			name: "testcase 2",
			args: args{
				root: NewBinaryTree(nil),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			PrintTree(tt.args.root)
		})
	}
}
