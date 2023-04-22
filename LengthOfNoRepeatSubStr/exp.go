package LengthOfNoRepeatSubStr

// check no repeat substring
func lengthOfLongestSubstring(s string) int {
	// abccba
	a := map[string]struct{}{}
	for _, ss := range s {
		if _, exist := a[string(ss)]; exist {

		}
	}

	return 0
}
