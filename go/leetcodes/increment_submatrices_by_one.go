package leetcodesgo

func rangeAddQueries(n int, queries [][]int) [][]int {
	prefixSum := make([][]int, n)
	for i := range prefixSum {
		prefixSum[i] = make([]int, n)
	}

	for _, query := range queries {
		row1, col1, row2, col2 := query[0], query[1], query[2], query[3]
		for row := row1; row <= row2; row++ {
			prefixSum[row][col1] += 1
			if col2 < n-1 {
				prefixSum[row][col2+1] -= 1
			}
		}
	}

	for i := range prefixSum {
		for j := 1; j < n; j++ {
			prefixSum[i][j] += prefixSum[i][j-1]
		}
	}

	return prefixSum
}
