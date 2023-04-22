package LongestCommonSubSequence

//
//func LongestCommonSubsequence(text1 string, text2 string) int {
//	// 最长公共子序列
//	// 使用dp[i][j] 表示 两个字符串的 text1[0,i] 和 text2[0,j] 的公共子序列
//	// 所以 dp[0][j] 和 dp[i][0] 都是0
//	// dp[i][j] 的值可以由dp[i-1][j-1]得到，但分两种情况
//	// 1：text[i] = text2[j] 则 dp[i][j] = dp[i-1][j-1] + 1
//	// 2:text[i] != text2[j] 则 dp[i][j] = max(dp[i-1][j], dp[i][j-1])
//	// 状态转换方程：
//	// dp[i][j] = dp[i-1][j-1] + 1  （text[i] = text2[j] ）
//	// dp[i][j] = max(dp[i-1][j], dp[i][j-1]) （text[i] != text2[j]）
//
//	dpList := make([][]int, len(text1)+1)
//	for index := range dpList{
//		dpList[index] = make([]int, len(text2)+1)
//	}
//	max := func(a,b int) int {
//		if a>b{
//			return a
//		}
//		return b
//	}
//	// 因为0已经定义了
//	for i:=1;i<len(text1)+1;i++{
//		for j:=1;j<len(text2)+1;j++{
//			// 子串的下标需要从0开始
//			if text1[i-1] == text2[j-1]{
//				dpList[i][j] = dpList[i-1][j-1] + 1
//			}else{
//				// 不一致的情况
//				dpList[i][j] = max(dpList[i-1][j], dpList[i][j-1])
//			}
//		}
//	}
//	return dpList[len(text1)][len(text2)]
//}

func LongestCommonSubsequence(text1 string, text2 string) int {
	// 最长公共子序列
	// 动态规划五步：
	// 1：确定动态数组下标的含义
	// 2：动态数组初始化的部分
	// 3：确定状态转换方程
	// 4：确定循环顺序
	// 5：打印动态数组
	// i，j 分别表示 text1 和text2 的下标
	// 状态转换方程
	// 1：若text1[i] == text2[j]
	// 则 dp[i][j] = dp[i-1][j-1] +1
	// 2: 若text1[i] != text2[j]
	// max(dp[i-1][j], text[i][i-1])
	if text1 == "" || text2 == "" {
		return 0
	}
	dpList := make([][]int, len(text1)+1)
	for index := range dpList {
		dpList[index] = make([]int, len(text2)+1)
	}

	for i := 0; i < len(text1); i++ {
		for j := 0; j < len(text2); j++ {
			if text1[i] == text2[j] {
				dpList[i+1][j+1] = dpList[i][j] + 1
			} else {
				dpList[i+1][j+1] = max(dpList[i][j+1], dpList[i+1][j])
			}
		}
	}
	return dpList[len(text1)][len(text2)]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
