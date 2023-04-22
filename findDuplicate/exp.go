package findDuplicate

func FindDuplicate(nums []int) int {
	// 快慢指针
	// 将整个数组理解为一个index -》nums[index] 的图
	// 所以问题就转换为了 环形链表 寻找环入口的问题
	slow, fast := nums[0], nums[nums[0]]
	for slow != fast {
		slow, fast = nums[slow], nums[nums[fast]]
	}
	slow = 0
	for nums[slow] != nums[fast] {
		slow = nums[slow]
		fast = nums[fast]
	}
	return nums[slow]
}
