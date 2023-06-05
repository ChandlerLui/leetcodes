package findPeakElements

func findPeakElement(nums []int) int {
	// 寻找峰值
	// 取中间的数
	// if mid-1 > mid 向左走，否则向右走
	left, right := 0, len(nums)-1
	for left < right {
		mid := int(uint(left+right) / 2)
		if nums[mid] > nums[mid+1] {
			right = mid
		} else {
			left = mid + 1
		}
	}
	return right
}
