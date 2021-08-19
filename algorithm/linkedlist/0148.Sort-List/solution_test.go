package sortList

import (
	"github.com/sevenger/gopher30days/algorithm"
	"reflect"
	"testing"
)

func Test_sortList(t *testing.T) {
	type args struct {
		head *ListNode
	}
	tests := []struct {
		name string
		args args
		want *ListNode
	}{
		{
			name: "testcase 1",
			args: args{
				head: algorithm.NewLinkedListByArgs(6, 5, 4, 3, 2, 1),
			},
			want: algorithm.NewLinkedListByArgs(1, 2, 3, 4, 5, 6),
		},
		{
			name: "testcase 2",
			args: args{
				head: algorithm.NewLinkedListByArgs(4, 2, 1, 3),
			},
			want: algorithm.NewLinkedListByArgs(1, 2, 3, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortList(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortList() = %v, want %v", algorithm.GetFormatLinkedListString(got), tt.want)
			}
		})
	}
}
