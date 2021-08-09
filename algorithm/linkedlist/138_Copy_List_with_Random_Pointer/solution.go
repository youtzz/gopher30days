package copyListWithRandomPointer

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

var cacheNode map[*Node]*Node

func init() {
	cacheNode = map[*Node]*Node{}
}

func copyRandomList(head *Node) *Node {
	if head == nil {
		return head
	}
	if n, has := cacheNode[head]; has {
		return n
	}

	newNode := &Node{Val: head.Val}
	cacheNode[head] = newNode
	newNode.Next = copyRandomList(head.Next)
	newNode.Random = copyRandomList(head.Random)
	return newNode
}
