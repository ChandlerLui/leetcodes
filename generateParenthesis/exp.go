package generateParenthesis

import (
	"strings"
)

// 括号生成
func generateParenthesis(n int) []string {
	// 回溯
	// 边界情况判断
	if n == 0 {
		return nil
	}
	if n == 1 {
		return []string{"()"}
	}
	ans := make([]string, 0)
	temp := make([]string, 0, n*2)
	var backTracking func(int, int)
	backTracking = func(left, right int) {
		if len(temp) == n*2 {
			ans = append(ans, strings.Join(temp, ""))
		}
		//
		if left < n {
			temp = append(temp, "(")
			backTracking(left+1, right)
			temp = temp[:len(temp)-1]

		}
		if right < left {
			temp = append(temp, ")")
			backTracking(left, right+1)
			temp = temp[:len(temp)-1]
		}
	}
	backTracking(0, 0)
	return ans
}
