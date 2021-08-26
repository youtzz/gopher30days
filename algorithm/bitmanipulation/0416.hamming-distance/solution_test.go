package problem0416

import "testing"

func Test_hammingDistance(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testcase 1",
			args: args{
				x: 1,
				y: 4,
			},
			want: 2,
		},
		{
			name: "testcase 2",
			args: args{
				x: 3,
				y: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingDistance(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("hammingDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hammingDistance2(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testcase 1",
			args: args{
				x: 1,
				y: 4,
			},
			want: 2,
		},
		{
			name: "testcase 2",
			args: args{
				x: 3,
				y: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingDistance2(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("hammingDistance2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hammingDistance_best(t *testing.T) {
	type args struct {
		x int
		y int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testcase 1",
			args: args{
				x: 1,
				y: 4,
			},
			want: 2,
		},
		{
			name: "testcase 2",
			args: args{
				x: 3,
				y: 1,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingDistance_best(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("hammingDistance_best() = %v, want %v", got, tt.want)
			}
		})
	}
}
