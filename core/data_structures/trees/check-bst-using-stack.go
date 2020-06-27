package trees

// Time complexity: O(N).
// Space complexity: O(N) to keep stack.
func isBSTUsingStack(root *node) bool {
	stack := []*node{}
	lastLeftNodeVal := -1 << 63

	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.left
		}

		root = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if root.data <= lastLeftNodeVal {
			return false
		}
		lastLeftNodeVal = root.data

		root = root.right
	}

	return true
}
