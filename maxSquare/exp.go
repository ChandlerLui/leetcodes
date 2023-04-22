package maxSquare

func maximalSquare(matrix [][]byte) int {
	// 最大正方形
	// 动态规划
	// 确定下标的含义， i,j 分别表示 横纵坐标，dp[j][j] 表示以ij为右下角的 最长的正方形的边
	// 确定递推公式 dp[i][j] = min(dp[i-1][j] + dp[i-1][j-1] + dp[i][j-1]) +1 (matrix[i][j] = 0)
	// 确定初始化的部分
	// 确定递推的顺序
	// 打印机结果
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	dp := make([][]int, len(matrix))
	var maxSide int
	// 初始化
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[0]))
		for j := 0; j < len(matrix[0]); j++ {
			dp[i][j] = int(matrix[i][j] - '0')
			if dp[i][j] == 1 {
				maxSide = 1
			}
		}
	}

	// 遍历
	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[0]); j++ {
			if dp[i][j] == 0 {
				continue
			}
			dp[i][j] = min(dp[i-1][j], min(dp[i-1][j-1], dp[i][j-1])) + dp[i][j]
			if dp[i][j] > maxSide {
				maxSide = dp[i][j]
			}
		}
	}
	return maxSide * maxSide
}
