package sort

type args struct {
	arr []int
}

var testCases = []struct {
	name string
	args args
	want []int
}{
	{
		name: "",
		args: args{
			arr: []int{7, 1, 2, 5},
		},
		want: []int{1, 2, 5, 7},
	},
	{
		name: "",
		args: args{
			arr: []int{0, 7, 1, 3, 8, 3, 1},
		},
		want: []int{0, 1, 1, 3, 3, 7, 8},
	},
	{
		name: "",
		args: args{
			arr: []int{0, 7, 1, 3, 8, 3, 1, -1},
		},
		want: []int{-1, 0, 1, 1, 3, 3, 7, 8},
	},
}
