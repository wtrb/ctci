package subarray

import "testing"

func Test_sumOfSubarray(t *testing.T) {
	type args struct {
		arr   []int
		start int
		end   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				arr:   []int{0, -2, 3, -2, 1},
				start: 1,
				end:   2,
			},
			want: 1,
		},
		{
			name: "",
			args: args{
				arr:   []int{0, -2, 3, -2, 1},
				start: 2,
				end:   2,
			},
			want: 3,
		},
		{
			name: "",
			args: args{
				arr:   []int{0, -2, 3, -2, 1},
				start: 2,
				end:   4,
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sumOfSubarray(tt.args.arr, tt.args.start, tt.args.end); got != tt.want {
				t.Errorf("sumOfSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
