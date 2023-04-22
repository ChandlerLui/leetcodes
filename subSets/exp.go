package subSets

func Subsets(nums []int) (ans [][]int) {
	// 子集
	ans = make([][]int, 0)
	set := make([]int, 0)

	var backTracking func(int2 int)
	backTracking = func(index int) {
		if index == len(nums) {
			ans = append(ans, append([]int{}, set...))
			return
		}
		// 考虑当前
		set = append(set, nums[index])
		backTracking(index + 1)
		set = set[:len(nums)-1]
		backTracking(index + 1)
	}
	backTracking(0)
	return
}
