package StringToInt

func MyAtoi(s string) int {
	// step0 check length of string
	// step1 get rid of space
	// step2 check first string is - or +
	// step3 range rest of str, add
	if len(s) > 2000 {
		return 0
	}
	m := map[string]int{"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9}
	var (
		isPositive  = true
		isFirstChar bool
	)
	var target int
	for _, r := range s {
		char := string(r)
		if char == " " {
			if isFirstChar {
				return 0
			}
			continue
		}
		if !isFirstChar && (char == "-" || char == "+") {
			if char == "-" {
				isPositive = false
			}
			isFirstChar = true
			continue
		}
		if !isFirstChar {
			// first Char
			isFirstChar = true
		}
		if res, exist := m[char]; exist {
			target = target*10 + res
		} else {
			break
		}
	}
	if !isPositive {
		if -target < -2147483648 {
			return -2147483648
		}
		return -target
	}
	if target > 2147483647 {
		return 2147483647
	}
	return target
}
