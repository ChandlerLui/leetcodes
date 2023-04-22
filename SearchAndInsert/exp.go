package SearchAndInsert

func SearchInsert(nums []int, target int) int {
	// 二分查找，寻找中间
	var (
		low  = 0
		high = len(nums) - 1
	)
	for low <= high {
		mid := low + (high-low)/2
		// 1,2 ,4,5  Tar:3
		if nums[mid] < target {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return low
}
