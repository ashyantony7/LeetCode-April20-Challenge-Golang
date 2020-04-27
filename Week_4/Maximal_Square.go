package main

func min(a, b byte) byte {
	if a < b {
		return a
	} else {
		return b
	}
}

func maximalSquare(matrix [][]byte) int {

	// get size of matrix
	m := len(matrix)
	if m == 0 {
		return 0
	}
	n := len(matrix[0])
	// maximum length of formed square
	var maxVal = int(matrix[0][0] - '0')

	// matrix to store intermediate results for
	// dynamic programming
	var DP = make([][]byte, m)
	for i := range DP {
		DP[i] = make([]byte, n)
		copy(DP[i], matrix[i])
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {

			// only do when the current val is 1
			if DP[i][j] == '1' {

				var temp byte = '0' // sq len as byte

				if (i > 0) && (j > 0) {
					temp = min(DP[i-1][j], DP[i][j-1])
					temp = min(temp, DP[i-1][j-1])
				}

				// get the square length
				sq_len := int(temp-'0') + 1
				// update the DP matrix
				DP[i][j] = byte(sq_len + '0')

				if sq_len > maxVal {
					maxVal = sq_len
				}
			}
		}
	}

	return maxVal * maxVal
}
