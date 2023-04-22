package conbineSUm

import "sort"

func CombinationSum(candidates []int, target int) (ans [][]int) {
	// 组合总数
	// 回溯
	if len(candidates) == 0 {
		return
	}
	comb := make([]int, 0)
	ans = make([][]int, 0)

	sort.Ints(candidates)

	var dfs func(target, start int)

	dfs = func(target, start int) {
		if target == 0 {
			// 已经满足
			tmp := make([]int, len(comb))
			copy(tmp, comb)
			ans = append(ans, append([]int{}, tmp...))
			return
		}
		if start > len(candidates)-1 {
			return
		}
		for i := start; i < len(candidates); i++ {
			if candidates[i] > target {
				break
			}
			comb = append(comb, candidates[i])
			dfs(target-candidates[i], i)
			comb = comb[:len(comb)-1]
		}
	}
	dfs(target, 0)
	return ans
}

// 主要在于递归中传递下一个数字
