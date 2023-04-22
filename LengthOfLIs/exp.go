package LengthOfLIs

//func LengthOfLIS(nums []int) int {
//	// 最长递增子序列
//	// 状态转换方程
//	// dp[i] 表示 以nums[i]为结尾的最长子序列的长度
//	// dp[i] 可以由dp[0,j] (j <= 1)中的 最大值 +1 得到（前提是小于0）
//	// 所以状态转移方程是：
//	// dp[i] = max(dp[j] + 1) (num[j] < nums[i], 0<=j<i)
//	if len(nums) == 0{
//		return 0
//	}
//	var max = 1
//	dpList := make([]int, len(nums), len(nums))
//	for i := 0; i < len(nums); i++ {
//		dpList[i] = 1 // 默认子序列长度是1
//		for j := 0; j < i; j++ {
//			if nums[j] < nums[i]{
//				// j的数值 小于 i的数值
//				if dpList[j] + 1 > dpList[i]{
//					dpList[i] = dpList[j] + 1
//				}
//				if dpList[i] > max{
//					max = dpList[i]
//				}
//			}
//		}
//	}
//	return max
//}

func LengthOfLIS(nums []int) int {
	// 最长递增子序列
	// i表示从0 -> i 的递增最大子序列的长度
	// 初始化：声明所有的item为1
	// 递推公式为：
	// dp[i] = max(dp[0,j])+1  if dp[i] > dp[j] (0=<j<i)
	// dp[i] = 1
	dpList := make([]int, len(nums), len(nums))
	var max int
	// [1,0,3,4,5]
	for i := 0; i < len(nums); i++ {
		dpList[i] = 1
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				if dpList[j]+1 > dpList[i] {
					dpList[i] = dpList[j] + 1
				}
			}
		}
		if dpList[i] > max {
			max = dpList[i]
		}
	}
	return max
}
