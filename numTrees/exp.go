package numTrees

func NumTrees(n int) int {
	// 动态规划
	// 以i为根的流量
	dp := make([]int, n+1)
	dp[0], dp[1] = 1, 1
	for i := 2; i <= n; i++ {
		// 从二开始，因为1，0 都已经OK了。
		for j := 1; j <= i; j++ {
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
