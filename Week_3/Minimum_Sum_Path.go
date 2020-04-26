package main

func min(a int, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func minPathSum(grid [][]int) int {

	// get the size of grid
	var m int = len(grid)
	var n int = len(grid[0])

	// make matrix to store min sum
	var minSum = make([][]int, m)
	for i := range minSum {
		minSum[i] = make([]int, n)
	}

	// loop through grid and get min sum
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i > 0 {
				if j > 0 {
					minSum[i][j] = min(minSum[i][j-1], minSum[i-1][j]) + grid[i][j]
				} else {
					minSum[i][j] = minSum[i-1][j] + grid[i][j]
				}
			} else {
				if j > 0 {
					minSum[i][j] = minSum[i][j-1] + grid[i][j]
				} else {
					minSum[i][j] = grid[i][j]
				}
			}
		}
	}

	return minSum[m-1][n-1]
}
