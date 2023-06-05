package searchRange

func SearchRange(nums []int, target int) []int {
	var binarySearch func(nums []int, target int) int

	binarySearch = func(nums []int, target int) int {
		left, right := 0, len(nums) //[)
		for left < right {
			mid := int(uint(left+right) / 2)
			if nums[mid] < target {
				left = mid + 1
			} else {
				right = mid
			}
		}
		return left
	}

	// first LTE
	start := binarySearch(nums, target)
	if start == len(nums) || nums[start] == target {
		return []int{-1, -1}
	}
	end := binarySearch(nums, target+1) - 1
	return []int{start, end}
}
