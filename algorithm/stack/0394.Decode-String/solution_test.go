package problem0394

import (
	"strings"
	"testing"
)

func Test_decodeString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "testcase 1",
			args: args{
				s: "3[a]2[b]",
			},
			want: "aaabb",
		},
		{
			name: "testcase 2",
			args: args{
				s: "3[a2[c]]",
			},
			want: "accaccacc",
		},
		{
			name: "testcase 3",
			args: args{
				s: "100[leetcode]",
			},
			want: func() string {
				sb := strings.Builder{}
				for i := 0; i < 100; i++ {
					sb.WriteString("leetcode")
				}
				return sb.String()
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeString(tt.args.s); got != tt.want {
				t.Errorf("decodeString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_decodeString_best(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "testcase 1",
			args: args{
				s: "3[a]2[b]",
			},
			want: "aaabb",
		},
		{
			name: "testcase 2",
			args: args{
				s: "3[a2[c]]",
			},
			want: "accaccacc",
		},
		{
			name: "testcase 3",
			args: args{
				s: "100[leetcode]",
			},
			want: func() string {
				sb := strings.Builder{}
				for i := 0; i < 100; i++ {
					sb.WriteString("leetcode")
				}
				return sb.String()
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := decodeString_best(tt.args.s); got != tt.want {
				t.Errorf("decodeString_best() = %v, want %v", got, tt.want)
			}
		})
	}
}
