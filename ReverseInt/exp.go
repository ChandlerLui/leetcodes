package ReverseInt

func reverse(x int) int {
	if x > -10 && x < 10 {
		return x
	}

	isPositive := x > 0
	var y = x
	if !isPositive {
		y = -x
	}
	var target int
	for y > 0 {
		target = target*10 + y%10
		y /= 10
	}
	if target > 2147483647 || -target < -2147483648 {
		return 0
	}
	if !isPositive {
		return -target
	}
	return target
}
