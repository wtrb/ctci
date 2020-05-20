package trees

import "testing"

func TestIsBST(t *testing.T) {
	bst := Tree{}
	bst.Insert(10)
	bst.Insert(5)
	bst.Insert(25)
	bst.Insert(1)
	bst.Insert(0)
	bst.Insert(19)
	bst.Insert(27)

	notBST := Tree{
		root: &node{
			data:  10,
			left:  &node{data: 20},
			right: &node{data: 1},
		},
	}

	type args struct {
		tree Tree
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "is bst",
			args: args{
				tree: bst,
			},
			want: true,
		},
		{
			name: "not bst",
			args: args{
				tree: notBST,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsBST(tt.args.tree); got != tt.want {
				t.Errorf("IsBST() = %v, want %v", got, tt.want)
			}
		})
	}
}
