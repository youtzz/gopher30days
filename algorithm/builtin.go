// Package algorithm 定义基础的数据结构类型，并与LeetCode中的保持一致
package algorithm

import (
	"fmt"
	"strconv"
	"strings"
)

// -----------------链表

type LinkList *ListNode

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewLinkList(arg []int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, v := range arg {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

func CompareLinkList(l1, l2 *ListNode) bool {
	rst := true
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			rst = false
			break
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	if l1 != nil || l2 != nil {
		rst = false
	}

	return rst
}

func PrintLinkList(head *ListNode) {
	fmt.Println(getFormatLinkListString(head))
}

func getFormatLinkListString(head *ListNode) string {
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

// -----------------树

type Tree *TreeNode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// -----------------矩阵

type Matrix [][]int
type Row []int

func NewMatrix(args ...[]int) [][]int {
	matrix := make([][]int, len(args))
	for i, arg := range args {
		matrix[i] = arg
		//for _, v := range arg {
		//	matrix[i] = append(matrix[i], v)
		//}
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
