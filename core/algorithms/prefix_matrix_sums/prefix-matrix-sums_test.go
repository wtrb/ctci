package submatrix

import (
	"reflect"
	"testing"
)

func Test_prefixSums(t *testing.T) {
	type args struct {
		mat [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "",
			args: args{
				mat: [][]int{
					{1, 2, 3},
					{4, 5, 6},
					{7, 8, 9},
				},
			},
			want: [][]int{
				{0, 0, 0, 0},
				{0, 1, 3, 6},
				{0, 5, 12, 21},
				{0, 12, 27, 45},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := prefixSums(tt.args.mat); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("prefixSums(%+v) = %v, want %v", tt.args.mat, got, tt.want)
			}
		})
	}
}

func Test_matrixBlockSum(t *testing.T) {
	type args struct {
		mat [][]int
		K   int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "",
			args: args{
				mat: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
				K:   1,
			},
			want: [][]int{{12, 21, 16}, {27, 45, 33}, {24, 39, 28}},
		},
		{
			name: "",
			args: args{
				mat: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
				K:   2,
			},
			want: [][]int{{45, 45, 45}, {45, 45, 45}, {45, 45, 45}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := matrixBlockSum(tt.args.mat, tt.args.K); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("matrixBlockSum(%+v, %v) = %v, want %v", tt.args.mat, tt.args.K, got, tt.want)
			}
		})
	}
}
