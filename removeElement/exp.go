package removeElement

func RemoveElement(nums []int, val int) int {

	// O1 原地去除 对应的val元素
	right := len(nums) - 1
	// 1,1,2,2,3,3
	for i := 0; i <= right; i++ {
		for ; i <= right && nums[i] == val; right-- {
			nums[i], nums[right] = nums[right], nums[i]
		}
	}
	return right + 1
}

func RemoveElement2(nums []int, val int) int {
	// 快慢指针的思路
	slow, fast := 0, 0
	for fast = 0; fast < len(nums); fast++ {
		if nums[fast] != val {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
