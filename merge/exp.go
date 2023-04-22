package merge

import (
	"sort"
)

// 合并多个重合的子区间
func merge(intervals [][]int) [][]int {
	if len(intervals) <= 1 {
		return intervals
	}

	// 按int[0]进行排序
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	pre := make([]int, 0)
	ans := make([][]int, 0)
	for i := 1; i < len(intervals); i++ {
		// 从1开始， 两种情况：
		// 1 如果interval[i][0] > pre[1],则两个不重合，直接更新
		// 2 如果interval[i][0] < pre[1],将pre[1]更新为最大值
		cur := intervals[i]
		if cur[0] > pre[1] {
			ans = append(ans, append([]int{}, pre...))
			pre = cur
		} else {
			pre[1] = max(pre[1], cur[1])
		}
	}
	ans = append(ans, pre)

	return ans
}
