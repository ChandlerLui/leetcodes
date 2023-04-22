package PalindromeNumber

import (
	"strconv"
)

// IsPalindrome 回文数字
func IsPalindrome(x int) bool {

	if x < 0 {
		return false
	}
	if x >= 0 && x < 10 {
		return true
	}
	xStr := strconv.Itoa(x)
	// range with string items, if str[len-1-i] equal with str[i]
	for i := 0; i < len(xStr); i++ {
		end := len(xStr) - 1 - i
		if i > end {
			return false
		}

		if xStr[i] == xStr[len(xStr)-1-i] {
			if end-i <= 1 {
				return true
			}
			continue
		} else {
			return false
		}
	}
	return false
}
