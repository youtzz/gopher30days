package Open_the_Lock

import "testing"

func Test_openLock(t *testing.T) {
	type args struct {
		deadends []string
		target   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testcase 1",
			args: args{
				[]string{"0201", "0101", "0102", "1212", "2002"}, "0202",
			},
			want: 6,
		},
		{
			name: "testcase 2",
			args: args{
				[]string{"8888"}, "0009",
			},
			want: 1,
		},
		{
			name: "testcase 3",
			args: args{
				[]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := openLock(tt.args.deadends, tt.args.target); got != tt.want {
				t.Errorf("openLock() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_openLock_double_bfs(t *testing.T) {
	type args struct {
		deadends []string
		target   string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "testcase 1",
			args: args{
				[]string{"0201", "0101", "0102", "1212", "2002"}, "0202",
			},
			want: 6,
		},
		{
			name: "testcase 2",
			args: args{
				[]string{"8888"}, "0009",
			},
			want: 1,
		},
		{
			name: "testcase 3",
			args: args{
				[]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := openLock_double_bfs(tt.args.deadends, tt.args.target); got != tt.want {
				t.Errorf("openLock_double_bfs() = %v, want %v", got, tt.want)
			}
		})
	}
}
