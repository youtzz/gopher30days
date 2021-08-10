package firstbadversion

import "testing"

type args struct {
	n int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{
		name: "testcase 1",
		args: args{
			n: 5,
		},
		want: 4,
	},
}

func Test_firstBadVersion(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstBadVersion(tt.args.n); got != tt.want {
				t.Errorf("firstBadVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_firstBadVersion_best(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := firstBadVersion_best(tt.args.n); got != tt.want {
				t.Errorf("firstBadVersion_best() = %v, want %v", got, tt.want)
			}
		})
	}
}
