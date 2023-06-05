package letterCombinations

import (
	"strconv"
	"strings"
)

// LetterCombinations 电话好吗的字母组合
func LetterCombinations(digits string) []string {
	//
	letters := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

	ans := make([]string, 0)
	combine := make([]string, len(digits))

	var backTracking func(index int)
	backTracking = func(index int) {
		if index == len(digits) {
			ans = append(ans, strings.Join(combine, ""))
			return
		}
		i, err := strconv.Atoi(string(digits[index]))
		if err != nil {
			return
		}

		letter := letters[i]
		for _, c := range letter {
			// 枚举第i个
			combine[index] = string(c)
			//combine = append(combine, string(c))
			backTracking(index + 1)
			//combine = combine[:len(combine)-1]
		}
	}
	backTracking(0)
	return ans
}
