package problem0026

import "testing"

func Test_removeDuplicates(t *testing.T) {
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
				nums: []int{1, 1, 2},
			},
			want: 2,
		},
		{
			name: "testcase 2",
			args: args{
				nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates(tt.args.nums)
			if got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
			// check if nums has duplicates
			set := make(map[int]int, got)
			for i := 0; i < got; i++ {
				val := tt.args.nums[i]
				set[val]++
				if set[val] > 1 {
					t.Errorf("removeDuplicates() do not remove the duplicates val %v", val)
				}
			}
		})
	}
}

func Test_removeDuplicates_best(t *testing.T) {
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
				nums: []int{1, 1, 2},
			},
			want: 2,
		},
		{
			name: "testcase 2",
			args: args{
				nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := removeDuplicates_best(tt.args.nums)
			if got != tt.want {
				t.Errorf("removeDuplicates_best() = %v, want %v", got, tt.want)
			}
			// check if nums has duplicates
			set := make(map[int]int, got)
			for i := 0; i < got; i++ {
				val := tt.args.nums[i]
				set[val]++
				if set[val] > 1 {
					t.Errorf("removeDuplicates_best() do not remove the duplicates val %v", val)
				}
			}
		})
	}
}
