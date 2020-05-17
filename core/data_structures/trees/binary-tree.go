package trees

import "fmt"

type Node interface {
	Insert(data int)
	Search(data int) bool
	PrintInOrder()
}

type node struct {
	data  int
	left  *node
	right *node
}

func New() Node {
	return &node{}
}

func (n *node) Insert(data int) {
	if data <= n.data {
		if n.left == nil {
			n.left = &node{data: data, left: nil, right: nil}

		} else {
			n.left.Insert(data)
		}

	} else {
		if n.right == nil {
			n.right = &node{data: data, left: nil, right: nil}

		} else {
			n.right.Insert(data)
		}
	}
}

func (n *node) Search(data int) bool {
	if data == n.data {
		return true

	} else if data < n.data {
		if n.left == nil {
			return false

		} else {
			return n.left.Search(data)
		}

	} else {
		if n.right == nil {
			return false

		} else {
			return n.right.Search(data)
		}
	}
}

func (n *node) PrintInOrder() {
	if n.left != nil {
		n.left.PrintInOrder()
	}
	fmt.Println(n.data)
	if n.right != nil {
		n.right.PrintInOrder()
	}
}
