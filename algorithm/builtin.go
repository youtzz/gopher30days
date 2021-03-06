// Package algorithm 定义基础的数据结构类型，并与LeetCode中的保持一致
package algorithm

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
	"strings"
)

const null = math.MinInt
const NULL = null

// ---------------------------------线性表---------------------------------

type LinerList = []int

func NewLinerList(args ...int) LinerList {
	s := make(LinerList, len(args))
	for i, v := range args {
		s[i] = v
	}
	return s
}

// ----------------------------------链表----------------------------------

type LinkedList = ListNode

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewLinkedList(args []int) *ListNode {
	return NewLinkedListByArgs(args...)
}

func NewLinkedListByArgs(args ...int) *ListNode {
	dummy := &ListNode{}
	cur := dummy
	for _, v := range args {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

// NewCycleLinedList 创建环形链表
func NewCycleLinedList(arr []int, pos int) *ListNode {
	if pos == -1 {
		return NewLinkedList(arr)
	}

	dummy := &ListNode{}
	prev := dummy
	var start, end *ListNode
	for i, v := range arr {
		node := &ListNode{Val: v}
		prev.Next = node
		prev = prev.Next
		if i == pos {
			start = node
		}
		if i == len(arr)-1 {
			end = node
		}
	}
	end.Next = start
	return dummy.Next
}

// CompareLinkedList 比较链表是否相同
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

// HasCycle 检查链表是否含有环
func HasCycle(head *ListNode) (*ListNode, bool) {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p, true
		}
	}
	return nil, false
}

// PrintLinkedList 格式化打印链表
func PrintLinkedList(head *ListNode) {
	fmt.Println(GetFormatLinkedListString(head))
}

// GetFormatLinkedListString 得到格式化链表字符串
func GetFormatLinkedListString(head *ListNode) string {
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

// JoinLinkedList 拼接两条链表
func JoinLinkedList(head *ListNode, next *ListNode) {
	for head != nil {
		if head.Next == nil {
			head.Next = next
			break
		}
		head = head.Next
	}
}

// -----------------------------------树-----------------------------------

type Tree = TreeNode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// NewBinaryTree 使用层先法创建二叉树
func NewBinaryTree(arr []int) *TreeNode {
	n := len(arr)
	if n == 0 {
		return nil
	}

	root := &TreeNode{
		Val: arr[0],
	}

	queue := make([]*TreeNode, 1, n*2)
	queue[0] = root

	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]

		if i < n && arr[i] != NULL {
			node.Left = &TreeNode{Val: arr[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < n && arr[i] != NULL {
			node.Right = &TreeNode{Val: arr[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

func NewBinaryTreeByArgs(args ...int) (root *TreeNode) {
	return NewBinaryTree(args)
}

func CompareBinaryTree(root1, root2 *TreeNode) bool {
	// 使用先序遍历+中序遍历确定两颗树相同
	return reflect.DeepEqual(PreOrder(root1), PreOrder(root2)) &&
		reflect.DeepEqual(InOrder(root1), InOrder(root2))

	// 其实直接DeepEqual就可以了
	//return reflect.DeepEqual(root1, root2)
}

// PrintTree 使用层先法打印树
func PrintTree(root *TreeNode) {
	if root == nil {
		fmt.Println("{ Empty Tree }")
		return
	}

	var res [][]string
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		isTail := true // 标志到树底
		var list []string

		l := len(queue)
		for i := 0; i < l; i++ {
			node := queue[0]
			queue = queue[1:]
			if node == nil {
				list = append(list, "#")
				queue = append(queue, nil, nil)
			} else {
				list = append(list, strconv.Itoa(node.Val))
				queue = append(queue, node.Left, node.Right)
				isTail = false // 只要还有节点就不是树底
			}
		}
		res = append(res, list)
		// 到达树底break
		if isTail {
			break
		}
	}

	// 处理的不是很好
	maxLen := len(res[len(res)-2]) / 2
	for i := 0; i < len(res)-1; i++ {
		for j := 0; j < maxLen-len(res[i])/2; j++ {
			fmt.Print("  ")
		}
		fmt.Println(res[i])
	}
}

// PreOrder 先序非递归遍历
func PreOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	var stack []*TreeNode
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

// PreOrderRecursive 先序递归遍历
func PreOrderRecursive(root *TreeNode, res *[]int) {
	if root == nil {
		return
	}
	*res = append(*res, root.Val)
	PreOrderRecursive(root.Left, res)
	PreOrderRecursive(root.Right, res)
}

// InOrder 中序非递归遍历
func InOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	var stack []*TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		root = node.Right
	}
	return res
}

// PostOrder 后序非递归遍历
func PostOrder(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	var stack []*TreeNode
	var lastVisit *TreeNode
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}

		node := stack[len(stack)-1]
		if node.Right == nil || node.Right == lastVisit {
			stack = stack[:len(stack)-1]
			res = append(res, node.Val)
			lastVisit = node
		} else {
			root = node.Right
		}
	}
	return res
}

// Dfs 深度优先遍历
func Dfs(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	var res []int
	PreOrderRecursive(root, &res)
	return res
}

//Dfs_divideAndConquer 深度优先遍历分治法
func Dfs_divideAndConquer(root *TreeNode) []int {
	return divideAndConquer(root)
}

func divideAndConquer(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var res []int
	left := divideAndConquer(root.Left)
	right := divideAndConquer(root.Right)

	res = append(res, root.Val)
	res = append(res, left...)
	res = append(res, right...)
	return res
}

// LevelOrder Bfs广度优先遍历
func LevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var res [][]int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		var list []int
		// 当前层的长度
		levelLen := len(queue)
		for i := 0; i < levelLen; i++ {
			// 出队
			node := queue[0]
			queue = queue[1:]
			list = append(list, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, list)
	}
	return res
}

// ----------------------------------矩阵----------------------------------

type Matrix = [][]int
type Row = []int

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

// ----------------------------------排序----------------------------------

// MergeSort 归并排序
func MergeSort(arr []int) []int {
	return mergeSort(arr)
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2
	left := mergeSort(arr[0:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	var res []int
	l, r := 0, 0

	for l < len(left) && r < len(right) {
		v1, v2 := left[l], right[r]
		if v1 < v2 {
			res = append(res, v1)
			l++
		} else {
			res = append(res, v2)
			r++
		}
	}
	res = append(res, left[l:]...)
	res = append(res, right[r:]...)
	return res
}

// QuickSort 快速排序
func QuickSort(arr []int) []int {
	quickSort(arr, 0, len(arr)-1)
	return arr
}

func quickSort(arr []int, low, high int) {
	if low < high {
		pivot := partition(arr, low, high)
		quickSort(arr, 0, pivot-1)
		quickSort(arr, pivot+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[low]
	for low < high {
		for low < high && arr[high] >= pivot {
			high--
		}
		arr[low] = arr[high]

		for low < high && arr[low] <= pivot {
			low++
		}
		arr[high] = arr[low]
	}
	arr[low] = pivot

	return low
}
