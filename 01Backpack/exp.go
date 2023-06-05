package _1Backpack

// 01 背包问题
type Item struct {
	Val    int
	Weight int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func solution(items []*Item, target int) int {
	// 从n个物品中选择最大的价值和
	var ans int
	var dfs func(int, int) int

	dfs = func(i int, subTarget int) int {
		if i < 0 {
			return 0
		}
		if items[i].Weight > subTarget {
			// 当前的条件是
			return dfs(i-1, subTarget)
		}
		return max(dfs(i-1, subTarget), dfs(i-1, subTarget-items[i].Weight)+items[i].Val)
	}

	return ans
}
