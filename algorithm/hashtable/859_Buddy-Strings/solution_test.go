package _859_Buddy_Strings

import "testing"

func Test_buddyStrings(t *testing.T) {
	type args struct {
		s    string
		goal string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "testcase 1",
			args: args{
				s:    "ab",
				goal: "ba",
			},
			want: true,
		},
		{
			name: "testcase 2",
			args: args{
				s:    "ab",
				goal: "ab",
			},
			want: false,
		},
		{
			name: "testcase 3",
			args: args{
				s:    "aa",
				goal: "aa",
			},
			want: true,
		},
		{
			name: "testcase 4",
			args: args{
				s:    "aaaaaaabc",
				goal: "aaaaaaacb",
			},
			want: true,
		},
		{
			name: "testcase 5",
			args: args{
				s:    "",
				goal: "aa",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buddyStrings(tt.args.s, tt.args.goal); got != tt.want {
				t.Errorf("buddyStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
