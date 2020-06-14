package wines

import "testing"

func TestMaxProfit(t *testing.T) {
	type args struct {
		prices []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "",
			args: args{
				prices: []int{1, 4, 2, 3},
			},
			want: 29,
		},
		{
			name: "",
			args: args{
				prices: []int{2, 3, 5, 1, 4},
			},
			want: 50,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MaxProfit(tt.args.prices); got != tt.want {
				t.Errorf("MaxProfit(%v) = %v, want %v", tt.args.prices, got, tt.want)
			}
		})
	}
}
