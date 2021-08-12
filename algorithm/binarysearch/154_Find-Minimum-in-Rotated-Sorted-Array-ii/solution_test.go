package _54_Find_Minimum_in_Rotated_Sorted_Array_ii

import (
	"github.com/sevenger/gopher30days/algorithm"
	"testing"
)

func Test_findMin(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testcase 1",
			args: args{
				nums: algorithm.NewSlice(1, 3, 5),
			},
			want: 1,
		},
		{
			name: "testcase 2",
			args: args{
				nums: algorithm.NewSlice(2, 2, 2, 0, 1),
			},
			want: 0,
		},
		{
			name: "testcase 3",
			args: args{
				nums: algorithm.NewSlice(10, 1, 10, 10, 10),
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.args.nums); got != tt.want {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
