package subSets

func Subsets(nums []int) (ans [][]int) {
	// 子集
	ans = make([][]int, 0)
	path := make([]int, 0)

	var backTracking func(index int)
	backTracking = func(index int) {
		ans = append(ans, append([]int{}, path...))
		if index == len(nums) {
			return
		}
		for i := index; i < len(nums); i++ {
			path = append(path, nums[i])
			backTracking(i + 1)
			path = path[:len(path)-1] //
		}
	}
	backTracking(0)
	return
}
