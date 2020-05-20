package bitmnp

import (
	"reflect"
	"testing"
)

func Test_generateSubsets(t *testing.T) {
	type args struct {
		char []rune
	}
	tests := []struct {
		name string
		args args
		want [][]rune
	}{
		{
			name: "a",
			args: args{
				char: []rune{'a'},
			},
			want: [][]rune{
				{' '},
				{'a'},
			},
		},
		{
			name: "ab",
			args: args{
				char: []rune{'a', 'b'},
			},
			want: [][]rune{
				{' ', ' '},
				{'a', ' '},
				{' ', 'b'},
				{'a', 'b'},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := generateSubsets(tt.args.char); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("generateSubsets() = %v, want %v", got, tt.want)
			}
		})
	}
}
