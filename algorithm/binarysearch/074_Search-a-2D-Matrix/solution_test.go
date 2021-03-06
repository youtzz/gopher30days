package searcha2dmatrix

import (
	. "github.com/sevenger/gopher30days/algorithm"
	"testing"
)

type args struct {
	matrix [][]int
	target int
}

var tests = []struct {
	name string
	args args
	want bool
}{
	{
		name: "testcase1",
		args: args{
			matrix: NewMatrix(NewRow(1, 3, 5, 7),
				NewRow(10, 11, 16, 20),
				NewRow(23, 30, 34, 60)),
			target: 3,
		},
		want: true,
	},
	{
		name: "testcase2",
		args: args{
			matrix: NewMatrix(NewRow(1, 3, 5, 7),
				NewRow(10, 11, 16, 20),
				NewRow(23, 30, 34, 60)),
			target: 13,
		},
		want: false,
	},
	{
		name: "testcase3",
		args: args{
			matrix: NewMatrix(NewRow(1)),
			target: 1,
		},
		want: true,
	},
}

func Test_searchMatrix(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("searchMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_searchMatrix_best(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix_best(tt.args.matrix, tt.args.target); got != tt.want {
				t.Errorf("searchMatrix_best() = %v, want %v", got, tt.want)
			}
		})
	}
}
