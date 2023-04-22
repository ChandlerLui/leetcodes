package MaxSubArray

// maxSubArray
// 动态规划
// 大问题的答案是由小问题的来一步一步推导来的
/*
状态转换方程
1：dp[i] := dp[i-1] + num[i]  (dp[i-1] >0)
2: dp[i] := nums[i] (dp[i-1 =<0])
*/

func MaxSubArray(nums []int) int {
	dpList := make([]int, len(nums))
	dpList[0] = nums[0] // 默认以第一个为结尾的就是第一个
	var max = dpList[0]
	for i := 1; i < len(nums); i++ {
		if dpList[i-1] > 0 {
			dpList[i] = dpList[i-1] + nums[i]
		} else {
			dpList[i] += nums[i]
		}
		if dpList[i] > max {
			max = dpList[i]
		}
	}
	return max
}
