package path

import "testing"

func Test_minCost(t *testing.T) {
	type args struct {
		cost [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				cost: [][]int{
					{1, 1, 1},
					{1, 1, 1},
					{1, 1, 1},
				},
			},
			want: 5,
		},
		{
			name: "",
			args: args{
				cost: [][]int{
					{1, 1, 1},
					{1, 3, 1},
					{1, 2, 1},
				},
			},
			want: 5,
		},
		{
			name: "",
			args: args{
				cost: [][]int{
					{1, 1, 3},
					{1, 3, 1},
					{1, 2, 1},
				},
			},
			want: 6,
		},
		{
			name: "",
			args: args{
				cost: [][]int{
					{1, 1, 3},
					{1, 3, 1},
					{1, 4, 1},
					{1, 1, 1},
				},
			},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minCost(tt.args.cost); got != tt.want {
				t.Errorf("minCost() = %v, want %v", got, tt.want)
			}
		})
	}
}
