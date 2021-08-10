package searchInsertPosition

import "testing"

func Test_searchInsert(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{nums: []int{1, 3, 5, 6}, target: 5}, 2},
		{"2", args{nums: []int{1, 3, 5, 6}, target: 2}, 1},
		{"3", args{nums: []int{1, 3, 5, 6}, target: 7}, 4},
		{"4", args{nums: []int{1, 3, 5, 6}, target: 0}, 0},
		{"5", args{nums: []int{1}, target: 0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchInsert() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchInsert2(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1", args{nums: []int{1, 3, 5, 6}, target: 5}, 2},
		{"2", args{nums: []int{1, 3, 5, 6}, target: 2}, 1},
		{"3", args{nums: []int{1, 3, 5, 6}, target: 7}, 4},
		{"4", args{nums: []int{1, 3, 5, 6}, target: 0}, 0},
		{"5", args{nums: []int{1}, target: 0}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchInsert2(tt.args.nums, tt.args.target); got != tt.want {
				t.Errorf("searchInsert2() = %v, want %v", got, tt.want)
			}
		})
	}
}
