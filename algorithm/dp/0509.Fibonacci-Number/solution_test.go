package problem0509

import "testing"

func Test_fib_dp(t *testing.T) {
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
				n: 0,
			},
			want: 0,
		},
		{
			name: "testcase 2",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "testcase 3",
			args: args{
				n: 30,
			},
			want: 832040,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib_dp(tt.args.n); got != tt.want {
				t.Errorf("fib_dp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fib_mem(t *testing.T) {
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
				n: 0,
			},
			want: 0,
		},
		{
			name: "testcase 2",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "testcase 3",
			args: args{
				n: 30,
			},
			want: 832040,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib_mem(tt.args.n); got != tt.want {
				t.Errorf("fib_mem() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_fib_recur(t *testing.T) {
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
				n: 0,
			},
			want: 0,
		},
		{
			name: "testcase 2",
			args: args{
				n: 1,
			},
			want: 1,
		},
		{
			name: "testcase 3",
			args: args{
				n: 30,
			},
			want: 832040,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fib_recur(tt.args.n); got != tt.want {
				t.Errorf("fib_recur() = %v, want %v", got, tt.want)
			}
		})
	}
}
