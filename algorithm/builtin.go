// Package algorithm 定义基础的数据结构类型，并与LeetCode中的保持一致
package algorithm

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
)

// ---------------------------------线性表---------------------------------

type Slice []int

func NewSlice(args ...int) Slice {
	s := make(Slice, len(args))
	for i, v := range args {
		s[i] = v
	}
	return s
}

// ----------------------------------链表----------------------------------

type LinkedList *ListNode

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewLinkedList(arg []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, v := range arg {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

func CompareLinkedList(l1, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	if l1 != nil || l2 != nil {
		return false
	}

	return true
}

func PrintLinkedList(head *ListNode) {
	fmt.Println(getFormatLinkedListString(head))
}

func getFormatLinkedListString(head *ListNode) string {
	var stack []int
	for head != nil {
		stack = append(stack, head.Val)
		head = head.Next
	}

	var sb strings.Builder
	sb.Grow(len(stack))
	sb.WriteString("{ ")
	for _, v := range stack {
		sb.WriteString(strconv.Itoa(v))
		sb.WriteString(" -> ")
	}
	sb.WriteString("nil }")
	return sb.String()
}

// -----------------------------------树-----------------------------------

type Tree *TreeNode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NewBinaryTree 使用层先法创建二叉树
func NewBinaryTree(args []int) *TreeNode {
	if args == nil || len(args) == 0 {
		return nil
	}
	nodes := make([]*TreeNode, len(args))
	for i, v := range args {
		var node *TreeNode
		// 传math.MaxInt64默认是空节点
		if v == math.MaxInt64 {
			node = nil
		} else {
			node = &TreeNode{Val: v}
		}
		nodes[i] = node
	}
	root, nodes := nodes[0], nodes[1:]

	stack := []*TreeNode{root}

	var node *TreeNode
	for len(stack) != 0 && len(nodes) != 0 {
		node, stack = stack[0], stack[1:]
		var left, right *TreeNode
		if len(nodes) != 0 {
			left = nodes[0]
			stack = append(stack, nodes[0])
			nodes = nodes[1:]
		}
		if len(nodes) != 0 {
			right = nodes[0]
			stack = append(stack, nodes[0])
			nodes = nodes[1:]
		}
		if node != nil {
			node.Left, node.Right = left, right
		}
	}
	return root
}

func CompareBinaryTree(root1, root2 *TreeNode) bool {
	// 使用先序遍历+中序遍历确定两颗树相同
	return reflect.DeepEqual(PreOrder(root1), PreOrder(root2)) &&
		reflect.DeepEqual(InOrder(root1), InOrder(root2))

	// 其实直接DeepEqual就可以了
	//return reflect.DeepEqual(root1, root2)
}

// PreOrder 先序非递归遍历
func PreOrder(root *TreeNode) []int {
	var stack []*TreeNode
	var res []int
	for root != nil || len(stack) > 0 {
		for root != nil {
			res = append(res, root.Val)
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		root = node.Right
	}
	return res
}

// InOrder 中序非递归遍历
func InOrder(root *TreeNode) []int {
	var stack []*TreeNode
	var res []int

	for root != nil && len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		res = append(res, node.Val)
		root = node.Right
	}
	return res
}

// PostOrder 后序非递归遍历
func PostOrder(root *TreeNode) []int {
	return nil
}

// PrintTree 使用层先法打印树
func PrintTree(root *TreeNode) {
	if root == nil {
		fmt.Println("{ Empty Tree }")
		return
	}
	stack := []*TreeNode{root}
	var tmpStack []*TreeNode
	var sb strings.Builder

	var node *TreeNode
	for len(stack) > 0 {
		for len(stack) > 0 {
			node, stack = stack[0], stack[1:]
			sb.WriteString(strconv.Itoa(node.Val))
			sb.WriteString(" ")
			if node.Left != nil {
				tmpStack = append(tmpStack, node.Left)
			}
			if node.Right != nil {
				tmpStack = append(tmpStack, node.Right)
			}
		}
		sb.WriteString("\n")
		stack = tmpStack
		tmpStack = tmpStack[len(tmpStack):]
	}
	fmt.Print(sb.String())
}

// ----------------------------------矩阵----------------------------------

type Matrix [][]int
type Row []int

func NewMatrix(args ...[]int) [][]int {
	matrix := make([][]int, len(args))
	for i, arg := range args {
		matrix[i] = arg
	}
	return matrix
}

func NewMatrixByRowLen(rowLen int, args ...int) [][]int {
	colCount := len(args) / rowLen
	matrix := make([][]int, colCount)

	for i := 0; i < colCount; i++ {
		for j := 0; j < rowLen; j++ {
			matrix[i] = append(matrix[i], args[j+(rowLen*i)])
		}
	}
	return matrix
}

func NewRow(args ...int) []int {
	row := make([]int, 0, len(args))
	for _, v := range args {
		row = append(row, v)
	}
	return row
}
