package trees

type Tree struct {
	root *node
}

func (t *Tree) Insert(data int) {
	if t.root == nil {
		t.root = &node{data: data}
		return
	}

	t.root.insert(data)
}

func (t *Tree) Contain(data int) bool {
	if t.root == nil {
		return false
	}

	return t.root.contain(data)
}

func (t *Tree) TraverseInOrder() []int {
	if t.root == nil {
		return []int{}
	}

	result := []int{}
	tranverseInOrder(t.root, &result)
	return result
}

func tranverseInOrder(n *node, out *[]int) {
	if n.left != nil {
		tranverseInOrder(n.left, out)
	}
	*out = append(*out, n.data)
	if n.right != nil {
		tranverseInOrder(n.right, out)
	}
}
