package trees

import (
	"reflect"
	"testing"
)

func TestTree_TraverseInOrder(t *testing.T) {
	bstTree := Tree{}
	bstTree.Insert(10)
	bstTree.Insert(5)
	bstTree.Insert(25)
	bstTree.Insert(1)
	bstTree.Insert(0)
	bstTree.Insert(19)
	bstTree.Insert(27)

	tests := []struct {
		name string
		tr   *Tree
		want []int
	}{
		{
			name: "",
			tr:   &bstTree,
			want: []int{0, 1, 5, 10, 19, 25, 27},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.tr.TraverseInOrder(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tree.TraverseInOrder() = %v, want %v", got, tt.want)
			}
		})
	}
}
