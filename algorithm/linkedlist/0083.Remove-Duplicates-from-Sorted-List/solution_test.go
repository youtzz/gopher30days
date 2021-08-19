package problem0083

import (
	. "github.com/sevenger/gopher30days/algorithm"
	"reflect"
	"testing"
)

func Test_deleteDuplicates(t *testing.T) {
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
				head: NewLinkedListByArgs(1, 1, 2),
			},
			want: NewLinkedListByArgs(1, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deleteDuplicates2(t *testing.T) {
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
				head: NewLinkedListByArgs(1, 1, 2),
			},
			want: NewLinkedListByArgs(1, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates2(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates2() = %v, want %v", GetFormatLinkedListString(got), tt.want)
			}
		})
	}
}

func Test_deleteDuplicates_recur(t *testing.T) {
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
				head: NewLinkedListByArgs(1, 1, 2),
			},
			want: NewLinkedListByArgs(1, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates_recur(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates_recur() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_deleteDuplicates_best(t *testing.T) {
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
				head: NewLinkedListByArgs(1, 1, 2),
			},
			want: NewLinkedListByArgs(1, 2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := deleteDuplicates_best(tt.args.head); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("deleteDuplicates_best() = %v, want %v", got, tt.want)
			}
		})
	}
}
