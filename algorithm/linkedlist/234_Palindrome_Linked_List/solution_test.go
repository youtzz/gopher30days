package palindromeLinkedList

import (
	"github.com/sevenger/gopher30days/algorithm"
	"testing"
)

func Test_isPalindrome_best(t *testing.T) {
	type args struct {
		head *algorithm.ListNode
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "",
			args: args{algorithm.NewLinkList([]int{1, 2, 2, 1})},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome_best(tt.args.head); got != tt.want {
				t.Errorf("isPalindrome_best() = %v, want %v", got, tt.want)
			}
		})
	}
}
