package trees

func IsBST(tree Tree) bool {
	// return isBST(tree.root, math.MinInt64, math.MaxInt64)
	// return isBST(tree.root, -1<<63, 1<<63-1)

	return isBSTUsingStack(tree.root)
}

// Time complexity: O(N) since we visit each node exactly once.
// Space complexity: O(N) since we keep up to the entire tree.
func isBST(node *node, min, max int) bool {
	if node == nil {
		return true
	}

	if node.data < min || node.data > max {
		return false
	}

	return isBST(node.left, min, node.data-1) && isBST(node.right, node.data+1, max)
}
