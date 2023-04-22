package sortColors

func SortColors(nums []int) {
	// 2,2,1,1,0,0
	p0, p2 := 0, len(nums)-1
	for i := 0; i < len(nums); i++ {
		for ; i < p2 && nums[i] == 2; p2-- {
			nums[i], nums[p2] = nums[p2], nums[i]
		}
		if nums[i] < 1 {
			nums[i], nums[p0] = nums[p0], nums[i]
			p0++
		}
	}
}
