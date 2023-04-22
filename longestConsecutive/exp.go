package longestConsecutive

//

func longestConsecutive(nums []int) int {
	// 先遍历将结果都存储在map中
	// 再遍历nums，判断是否是连续序列的第一个，并记录最大值
	existMap := make(map[int]bool)
	for _, num := range nums {
		existMap[num] = true
	}
	var max int
	for _, num := range nums {
		if existMap[num-1] {
			continue
		}
		// 不存在 则往后递增查找
		tmp := num
		var count int
		for existMap[tmp] {
			count++
			tmp++
		}
		if count > max {
			max = count
		}
	}
	return max
}
