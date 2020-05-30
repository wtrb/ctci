package search

import "testing"

func Test_binarySearch(t *testing.T) {
	type args struct {
		sortedArray []int
		x           int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "found at middle",
			args: args{
				sortedArray: []int{-3, 1, 3, 7, 12, 99, 101},
				x:           7,
			},
			want: true,
		},
		{
			name: "not found",
			args: args{
				sortedArray: []int{-3, 1, 3, 7, 12, 99, 101},
				x:           0,
			},
			want: false,
		},
		{
			name: "found on left side",
			args: args{
				sortedArray: []int{-3, 1, 3, 7, 12, 99, 101},
				x:           -3,
			},
			want: true,
		},
		{
			name: "found on right side",
			args: args{
				sortedArray: []int{-3, 1, 3, 7, 12, 99, 101},
				x:           12,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := binarySearch(tt.args.sortedArray, tt.args.x); got != tt.want {
				t.Errorf("binarySearch() = %v, want %v", got, tt.want)
			}
		})
	}
}
