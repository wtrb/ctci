package subarray

import "testing"

func Test_maxSubarraySum(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				arr: []int{-2, 2, 5, -11, 6},
			},
			want: 7,
		},
		{
			name: "",
			args: args{
				arr: []int{3, 2, -6, 4, 0},
			},
			want: 5,
		},
		{
			name: "",
			args: args{
				arr: []int{-10},
			},
			want: -10,
		},
		{
			name: "",
			args: args{
				arr: []int{-10, -3},
			},
			want: -3,
		},
		{
			name: "",
			args: args{
				arr: []int{-2, 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxSubarraySum(tt.args.arr); got != tt.want {
				t.Errorf("maxSubarraySum(%v) = %v, want %v", tt.args.arr, got, tt.want)
			}
		})
	}
}
