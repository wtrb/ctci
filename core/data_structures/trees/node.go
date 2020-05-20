package trees

import "fmt"

type node struct {
	data  int
	left  *node
	right *node
}

func (n *node) insert(data int) {
	if data <= n.data {
		if n.left == nil {
			n.left = &node{data: data, left: nil, right: nil}

		} else {
			n.left.insert(data)
		}

	} else {
		if n.right == nil {
			n.right = &node{data: data, left: nil, right: nil}

		} else {
			n.right.insert(data)
		}
	}
}

func (n *node) contain(data int) bool {
	if data == n.data {
		return true

	} else if data < n.data {
		if n.left == nil {
			return false

		} else {
			return n.left.contain(data)
		}

	} else {
		if n.right == nil {
			return false

		} else {
			return n.right.contain(data)
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
