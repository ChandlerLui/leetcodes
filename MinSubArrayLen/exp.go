package MinSubArrayLen

import "math"

func minSubArrayLen(s int, nums []int) int {
	// 最短的连续子序列 大于 s的
	var min = math.MaxInt64

	// 滑动窗口
	// fast slow 从0开始
	fast, slow := 0, 0
	for fast < len(nums) {
		if fast == slow && nums[fast] < s {
			fast++
		}
	}

	return min
}
