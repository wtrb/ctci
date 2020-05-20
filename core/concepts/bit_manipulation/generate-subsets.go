package bitmnp

func generateSubsets(char []rune) [][]rune {
	var n uint
	var i uint
	var j uint

	n = uint(len(char))

	result := make([][]rune, 1<<n)
	for i = 0; i < (1 << n); i++ {
		result[i] = []rune{}
		for j = 0; j < n; j++ {
			if (i & (1 << j)) != 0 {
				result[i] = append(result[i], char[j])

			} else {
				result[i] = append(result[i], ' ')
			}
		}
	}

	return result
}
