package leetcode

func minDistance(word1 string, word2 string) int {
	n := len(word1)
	m := len(word2)
	if len(word1) == 0 || len(word2) == 0 {
		return n + m
	}

	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}

	for i := 1; i < n+1; i++ {
		dp[i][0] = i
	}

	for j := 1; j < m+1; j++ {
		dp[0][j] = j
	}
	for i := 1; i < n+1; i++ {
		for j := 1; j < m+1; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i][j-1], dp[i-1][j], dp[i-1][j-1]) + 1
			}

		}
	}
	return dp[n][m]

}

func min(n ...int) int {
	m := n[0]
	for _, i := range n {
		if i < m {
			m = i
		}
	}
	return m
}
