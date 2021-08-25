package problem0160

import (
	"github.com/sevenger/gopher30days/algorithm"
	"reflect"
	"testing"
)

var headA, headB, public *ListNode
var usecaseNo int

func getUseCase(no int) (*ListNode, *ListNode, *ListNode) {
	usecase1 := func() (*ListNode, *ListNode, *ListNode) {
		list1 := algorithm.NewLinkedListByArgs(4, 1)
		list2 := algorithm.NewLinkedListByArgs(5, 0, 1)
		public := algorithm.NewLinkedListByArgs(8, 4, 5)
		algorithm.JoinLinkedList(list1, public)
		algorithm.JoinLinkedList(list2, public)
		return list1, list2, public
	}
	usecase2 := func() (*ListNode, *ListNode, *ListNode) {
		list1 := algorithm.NewLinkedListByArgs(0, 9, 1)
		list2 := algorithm.NewLinkedListByArgs(3)
		public := algorithm.NewLinkedListByArgs(2, 4)
		algorithm.JoinLinkedList(list1, public)
		algorithm.JoinLinkedList(list2, public)
		return list1, list2, public
	}
	usecase3 := func() (*ListNode, *ListNode, *ListNode) {
		public := algorithm.NewLinkedListByArgs()
		list1 := algorithm.NewLinkedListByArgs(2, 6, 4)
		list2 := algorithm.NewLinkedListByArgs(1, 5)
		algorithm.JoinLinkedList(list1, public)
		algorithm.JoinLinkedList(list2, public)
		return list1, list2, public
	}

	if usecaseNo == no {
		return headA, headB, public
	}

	usecaseNo = no

	switch no {
	case 1:
		headA, headB, public = usecase1()
	case 2:
		headA, headB, public = usecase2()
	case 3:
		headA, headB, public = usecase3()
	}
	return headA, headB, public
}

func getUseCaseList(useCaseNo, which int) *ListNode {
	headA, headB, public := getUseCase(usecaseNo)
	switch which {
	case 1:
		return headA
	case 2:
		return headB
	case 3:
		return public
	}
	return nil
}

func Test_getIntersectionNode(t *testing.T) {
	type args struct {
		headA *ListNode
		headB *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "testcase 1",
			args: args{
				headA: getUseCaseList(1, 1),
				headB: getUseCaseList(1, 2),
			},
			want: getUseCaseList(1, 3),
		},
		{
			name: "testcase 2",
			args: args{
				headA: getUseCaseList(2, 1),
				headB: getUseCaseList(2, 2),
			},
			want: getUseCaseList(2, 3),
		},
		{
			name: "testcase 3",
			args: args{
				headA: getUseCaseList(3, 1),
				headB: getUseCaseList(3, 2),
			},
			want: getUseCaseList(3, 3),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNode(tt.args.headA, tt.args.headB); !algorithm.CompareLinkedList(got, tt.want) {
				algorithm.PrintLinkedList(tt.args.headA)
				algorithm.PrintLinkedList(tt.args.headB)
				algorithm.PrintLinkedList(got)
				algorithm.PrintLinkedList(tt.want)
				t.Errorf("getIntersectionNode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getIntersectionNode2(t *testing.T) {
	type args struct {
		headA *ListNode
		headB *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "testcase 1",
			args: args{
				headA: getUseCaseList(1, 1),
				headB: getUseCaseList(1, 2),
			},
			want: getUseCaseList(1, 3),
		},
		{
			name: "testcase 2",
			args: args{
				headA: getUseCaseList(2, 1),
				headB: getUseCaseList(2, 2),
			},
			want: getUseCaseList(2, 3),
		},
		{
			name: "testcase 3",
			args: args{
				headA: getUseCaseList(3, 1),
				headB: getUseCaseList(3, 2),
			},
			want: getUseCaseList(3, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNode2(tt.args.headA, tt.args.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNode2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getIntersectionNode_best(t *testing.T) {
	type args struct {
		headA *ListNode
		headB *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "testcase 1",
			args: args{
				headA: getUseCaseList(1, 1),
				headB: getUseCaseList(1, 2),
			},
			want: getUseCaseList(1, 3),
		},
		{
			name: "testcase 2",
			args: args{
				headA: getUseCaseList(2, 1),
				headB: getUseCaseList(2, 2),
			},
			want: getUseCaseList(2, 3),
		},
		{
			name: "testcase 3",
			args: args{
				headA: getUseCaseList(3, 1),
				headB: getUseCaseList(3, 2),
			},
			want: getUseCaseList(3, 3),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getIntersectionNode_best(tt.args.headA, tt.args.headB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getIntersectionNode_best() = %v, want %v", got, tt.want)
			}
		})
	}
}
