package problem0229

import (
	"reflect"
	"testing"
)

func Test_majorityElement(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "testcase 1",
			args: args{
				nums: []int{3, 2, 3},
			},
			want: []int{3},
		},
		{
			name: "testcase 2",
			args: args{
				nums: []int{1},
			},
			want: []int{1},
		},
		{
			name: "testcase 3",
			args: args{
				nums: []int{1, 1, 1, 3, 3, 2, 2, 2},
			},
			want: []int{1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElement(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("majorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
