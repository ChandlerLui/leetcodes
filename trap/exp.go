package trap

import (
	"fmt"
)

func Trap(height []int) int {
	// 接雨水
	// 双数组的方法
	preMax, sufMax := make([]int, len(height)), make([]int, len(height))

	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	preMax[0] = height[0]
	for i := 1; i < len(height); i++ {
		preMax[i] = max(preMax[i-1], height[i])
	}
	sufMax[len(height)-1] = height[len(height)-1]
	for i := len(height) - 2; i >= 0; i-- {
		sufMax[i] = max(sufMax[i+1], height[i])
	}

	var ans int
	for i := 0; i <= len(height)-1; i++ {
		pre, suf, h := preMax[i], sufMax[i], height[i]
		ans += min(pre, suf) - h
	}
	fmt.Println(ans)
	return ans
}

func Trap2(height []int) int {
	// 双指针的方法
	max := func(a, b int) int {
		if a < b {
			return b
		}
		return a
	}

	left, right, ans := 0, len(height)-1, 0
	var leftMax, rightMax int
	for left < right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])

		if height[left] < height[right] {
			ans += leftMax - height[left]
			left++
		} else {
			ans += rightMax - height[right]
			right--
		}
	}
	return ans
}
