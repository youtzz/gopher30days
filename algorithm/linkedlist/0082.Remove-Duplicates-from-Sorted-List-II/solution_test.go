package problem0082

import (
	"github.com/sevenger/gopher30days/algorithm"
	"reflect"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
	type args struct {
		head *algorithm.ListNode
	}
	tests := []struct {
		name string
		args args
		want *algorithm.ListNode
	}{
		{
			name: "testcase 1",
			args: args{
				head: algorithm.NewLinkedListByArgs(1, 2, 2, 3, 3, 4),
			},
			want: algorithm.NewLinkedListByArgs(1, 4),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v",
					algorithm.GetFormatLinkedListString(got), algorithm.GetFormatLinkedListString(tt.want))
			}
		})
	}
}
