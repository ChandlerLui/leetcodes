package Permute

func Permute(nums []int) [][]int {
	// 全排列
	ans := make([][]int, 0)
	existMap := make(map[int]bool)
	temp := make([]int, 0)

	var backTracing func(int)
	backTracing = func(index int) {
		if len(temp) == len(nums) {
			ans = append(ans, append([]int{}, temp...))
			return
		}
		for i := 0; i < len(nums); i++ {
			if !existMap[nums[i]] {
				existMap[nums[i]] = true
				temp = append(temp, nums[i])
				backTracing(i + 1)
				temp = temp[:len(temp)-1]
				delete(existMap, nums[i])
			}
		}

	}
	backTracing(0)
	return ans
}
