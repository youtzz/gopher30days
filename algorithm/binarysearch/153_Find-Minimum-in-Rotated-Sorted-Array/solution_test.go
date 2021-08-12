package _53_Find_Minimum_in_Rotated_Sorted_Array

import "testing"

type args struct {
	nums []int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{
		name: "testcase 1",
		args: args{
			nums: []int{3, 4, 5, 1, 2},
		},
		want: 1,
	},
	{
		name: "testcase 2",
		args: args{
			nums: []int{4, 5, 6, 7, 0, 1, 2},
		},
		want: 0,
	},
	{
		name: "testcase 3",
		args: args{
			nums: []int{11, 13, 15, 17},
		},
		want: 11,
	},
}

func Test_findMin(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findMin(tt.args.nums); got != tt.want {
				t.Errorf("findMin() = %v, want %v", got, tt.want)
			}
		})
	}
}
