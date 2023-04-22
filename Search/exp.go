package Search

func Search(nums []int, target int) int {
	// 搜索旋转排序数组
	// 二分
	// 3，4，5，6, 7, 1，2
	low, high := 0, len(nums)-1
	// 1: 如果 low～mid 是有序区间 mid > target > low, 则 high 左移， 否则 low 右移
	// 2: 如果 mid ～ high 是有序区间，mid< target <high 则low右移，否则high左移
	for low <= high {
		mid := int(uint(low+high) / 2)
		if nums[mid] == target {
			return mid
		}
		if nums[mid] >= nums[low] {
			if nums[mid] > target && nums[low] <= target {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if nums[mid] < target && nums[high] >= target {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}

func Search2(nums []int, target int) int {
	// 搜索旋转排序数组
	// 如果low，mid是排序数组，即 low<target<mid, 则high左移，否则low右移
	// 如果mid，high的排序，即mid<target<high, 则low右移, 否则high左移
	low, high := 0, len(nums)-1
	for low <= high {
		mid := int(uint(low+high) / 2)
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > nums[low] {
			if nums[low] <= target && nums[mid] > target {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if nums[mid] < target && nums[high] >= target {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	return -1
}
