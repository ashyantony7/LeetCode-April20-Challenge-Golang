package main

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func longestCommonSubsequence(text1 string, text2 string) int {

	// make a subarray to store result for dynamic programming
	var len1, len2 int = len(text1), len(text2)
	var DP = make([][]int, len1+1)
	for i := range DP {
		DP[i] = make([]int, len2+1)

	}

	// store result for common length in text1[0..i] text2[0..j]
	for i := 0; i < len1; i++ {
		for j := 0; j < len2; j++ {
			if text1[i] == text2[j] {
				DP[i+1][j+1] = DP[i][j] + 1
			} else {
				DP[i+1][j+1] = max(DP[i][j+1], DP[i+1][j])
			}
		}
	}

	return DP[len1][len2]
}
