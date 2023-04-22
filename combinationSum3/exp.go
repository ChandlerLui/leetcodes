package combinationSum3

import "fmt"

// 找出所有相加之和为 n 的 k 个数的组合，且满足下列条件：
// 只使用数字1到9
// 每个数字 最多使用一次
func CombinationSum3(k int, n int) [][]int {
	ans := make([][]int, 0)
	comb := make([]int, 0)

	var backTracking func(int, int)
	backTracking = func(cur, target int) {
		// 满足条件
		if len(comb) == k && target == 0 {
			fmt.Println(comb)
			ans = append(ans, append([]int{}, comb...))
			return
		}
		if cur >= 10 || target < 0 || len(comb) > k {
			return
		}
		// 不选当前,target不变
		backTracking(cur+1, target)
		// 选当前
		comb = append(comb, cur)
		backTracking(cur+1, target-cur)
		comb = comb[:len(comb)-1]
	}
	backTracking(1, n)
	return ans
}
