package recursion

import "fmt"

func nQueens(board *[][]int, queens int) ([][]int, bool) {
	fmt.Printf("%+v", *board)
	fmt.Println()
	fmt.Printf("%+v", len(*board))
	fmt.Println()
	fmt.Printf("%v", queens)
	fmt.Println()

	if queens == 0 {
		return *board, true
	}

	for i := 0; i < len(*board); i++ {
		for j := 0; j < len((*board)[i]); j++ {
			if ok := isAttached(board, i, j); ok {
				// fmt.Printf("attched at [%v,%v]", i, j)
				// fmt.Println()
				continue

			} else {
				(*board)[i][j] = 1
			}

			if _, ok := nQueens(board, queens-1); ok {
				return *board, true

			} else {
				(*board)[i][j] = 0
			}
		}
	}

	return *board, false
}

func isAttached(board *[][]int, x, y int) bool {
	// fmt.Printf("%+v", *board)
	// fmt.Println()
	for j := 0; j < len((*board)[x]); j++ {
		fmt.Printf("checking at [%v,%v] = %v", x, j, (*board)[x][j])
		if (*board)[x][j] == 1 {
			// fmt.Printf("attched at [%v,%v]", x, j)
			return true
		}
	}
	for i := 0; i < len(*board); i++ {
		if (*board)[i][y] == 1 {
			// fmt.Printf("attched at [%v,%v]", i, y)
			return true
		}
	}
	for i := 0; i < len(*board); i++ {
		for j := 0; j < len((*board)[i]); j++ {
			if i+j == x+y || i-j == x-y {
				// fmt.Printf("attched at [%v,%v]", i, j)
				return true
			}
		}
	}

	return false
}
