package path

import "testing"

func Test_ways(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				matrix: [][]int{
					{0, 1, 0},
					{0, 1, 0},
					{1, 0, 0},
				},
			},
			want: 0,
		},
		{
			name: "",
			args: args{
				matrix: [][]int{
					{0, 1, 0},
					{0, 1, 1},
					{0, 0, 0},
				},
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				matrix: [][]int{
					{0, 1, 0},
					{0, 0, 1},
					{0, 0, 0},
				},
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				matrix: [][]int{
					{0, 1, 0},
					{0, 0, 1},
					{0, 0, 0},
					{0, 0, 0},
				},
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ways(tt.args.matrix); got != tt.want {
				t.Errorf("ways() = %v, want %v", got, tt.want)
			}
		})
	}
}
