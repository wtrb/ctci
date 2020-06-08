package recursion

func nQueens(board [][]int, queens int) bool {
	if queens == 0 {
		return true
	}

	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if ok := isAttached(board, i, j); ok {
				continue

			} else {
				board[i][j] = 1
			}

			if ok := nQueens(board, queens-1); ok {
				return true

			} else {
				board[i][j] = 0
			}
		}
	}

	return false
}

func isAttached(board [][]int, x, y int) bool {
	for j := 0; j < len(board[x]); j++ {
		if board[x][j] == 1 {
			return true
		}
	}
	for i := 0; i < len(board); i++ {
		if board[i][y] == 1 {
			return true
		}
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if (i+j == x+y || i-j == x-y) && board[i][j] == 1 {
				return true
			}
		}
	}

	return false
}
