package recursion

import (
	"reflect"
	"testing"
)

func Test_nQueens(t *testing.T) {
	type args struct {
		board  [][]int
		queens int
	}
	tests := []struct {
		name  string
		args  args
		want  [][]int
		want1 bool
	}{
		{
			name: "3 queens",
			args: args{
				board:  [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
				queens: 3,
			},
			want:  [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
			want1: false,
		},
		{
			name: "4 queens",
			args: args{
				board:  [][]int{{0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}},
				queens: 4,
			},
			want:  [][]int{{0, 1, 0, 0}, {0, 0, 0, 1}, {1, 0, 0, 0}, {0, 0, 1, 0}},
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := nQueens(&tt.args.board, tt.args.queens)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("nQueens() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("nQueens() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
