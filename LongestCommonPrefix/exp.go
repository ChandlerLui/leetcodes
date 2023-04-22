package LongestCommonPrefix

import "strings"

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}
	start := strs[0]
	for _, ss := range strs[1:] {
		for !strings.HasPrefix(ss, start) {
			start = start[:len(start)-1]
		}

	}
	return start
}
