package problem0070

import "testing"

func Test_climbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testcase 1",
			args: args{
				n: 3,
			},
			want: 3,
		},
		{
			name: "testcase 2",
			args: args{
				n: 5,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.args.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_climbStairs_best(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testcase 1",
			args: args{
				n: 3,
			},
			want: 3,
		},
		{
			name: "testcase 2",
			args: args{
				n: 5,
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs_best(tt.args.n); got != tt.want {
				t.Errorf("climbStairs_best() = %v, want %v", got, tt.want)
			}
		})
	}
}
