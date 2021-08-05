package algorithm

import (
	"reflect"
	"testing"
)

func TestCompareLinkList(t *testing.T) {
	type args struct {
		l1 *ListNode
		l2 *ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"equal", args{l1: NewLinkList([]int{1, 2, 3}), l2: NewLinkList([]int{1, 2, 3})}, true},
		{"equal", args{l1: nil, l2: nil}, true},
		{"not equal", args{l1: NewLinkList([]int{1, 2, 3, 4}), l2: NewLinkList([]int{1, 2, 3})}, false},
		{"not equal", args{l1: NewLinkList([]int{1, 2, 3, 4}), l2: nil}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareLinkList(tt.args.l1, tt.args.l2); got != tt.want {
				t.Errorf("CompareLinkList() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewLinkList(t *testing.T) {
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
			if got := NewLinkList(tt.args.arg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLinkList() = %v, want %v", got, tt.want)
			}
		})
	}
}
