package trees

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

func (n *node) height() int {
	lHeight := 0
	rHeight := 0

	if n.left != nil {
		lHeight = n.left.height()
	}
	if n.right != nil {
		rHeight = n.right.height()
	}

	if lHeight > rHeight {
		return lHeight + 1

	} else {
		return rHeight + 1
	}
}

// func (n *node) PrintInOrder() {
// 	if n.left != nil {
// 		n.left.PrintInOrder()
// 	}
// 	fmt.Println(n.data)
// 	if n.right != nil {
// 		n.right.PrintInOrder()
// 	}
// }
