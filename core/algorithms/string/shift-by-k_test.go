package string

import "testing"

func Test_shiftByK(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "hacker>>2",
			args: args{
				s: "hacker",
				k: 2,
			},
			want: "erhack",
		},
		{
			name: "hacker>>6",
			args: args{
				s: "hacker",
				k: 6,
			},
			want: "hacker",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shiftByK(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("shiftByK() = %v, want %v", got, tt.want)
			}
		})
	}
}
