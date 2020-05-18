package trees

func CheckBST(tree Tree) bool {
	return checkBST(tree.root, 1>>32, 1<<32)
}

func checkBST(node *node, min, max int) bool {
	if node == nil {
		return true
	}

	if node.data < min || node.data > max {
		return false
	}

	return checkBST(node.left, min, node.data-1) && checkBST(node.right, node.data+1, max)
}
