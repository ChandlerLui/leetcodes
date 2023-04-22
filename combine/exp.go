package combine

import "fmt"

// 生成，1～n 的 k为长度的
func Combine(n int, k int) [][]int {
	ans := make([][]int, 0)
	comb := make([]int, 0, k)

	var backtracking func(int)
	backtracking = func(i int) {
		if len(comb) == k {
			// done
			fmt.Println(comb)
			ans = append(ans, append([]int{}, comb...))
			return
		}
		if i > n {
			return
		}
		comb = append(comb, i)
		backtracking(i + 1)
		comb = comb[:len(comb)-1]
		backtracking(i + 1)
	}
	backtracking(1)
	return ans
}
