package bitmnp

import "testing"

func TestLargestPowerOfTwo16(t *testing.T) {
	type args struct {
		n int16
	}
	tests := []struct {
		name string
		args args
		want int16
	}{
		{
			name: "",
			args: args{
				n: 2,
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				n: 3,
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				n: 15,
			},
			want: 8,
		},
		{
			name: "",
			args: args{
				n: 258,
			},
			want: 256,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LargestPowerOfTwo16(tt.args.n); got != tt.want {
				t.Errorf("LargestPowerOfTwo16() = %v, want %v", got, tt.want)
			}
		})
	}
}
