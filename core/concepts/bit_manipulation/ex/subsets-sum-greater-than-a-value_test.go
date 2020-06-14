package bitmnp

import "testing"

func Test_countSubsets(t *testing.T) {
	type args struct {
		set []int
		val int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				set: []int{1, 2, 3},
				val: 5,
			},
			want: 2,
		},
		{
			name: "",
			args: args{
				set: []int{0, 1, 2, 3},
				val: 5,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSubsets(tt.args.set, tt.args.val); got != tt.want {
				t.Errorf("countSubsets(%v, %v) = %v, want %v", tt.args.set, tt.args.val, got, tt.want)
			}
		})
	}
}
