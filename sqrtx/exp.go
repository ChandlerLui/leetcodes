package sqrtx

func mySqrt(x int) int {
	// 获取x的中位数，二分终止条件是low《= high
	var (
		high = x
		low  = 0
		res  = 0
	)
	for low <= high {
		mid := low + (high-low)/2
		if mid*mid <= x {
			res = mid
			mid += 1
		} else {
			mid -= 1
		}
	}
	return res
}
