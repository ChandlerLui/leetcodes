package threeSum

import (
	"sort"
)

func ThreeSum(nums []int) [][]int {
	res := make([][]int, 0)
	if len(nums) < 3 {
		return res
	}
	// 排序
	sort.Ints(nums)

	for first := 0; first < len(nums); first++ {
		if nums[first] > 0 {
			break
		}
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		left := first + 1
		right := len(nums) - 1
		for left < right {
			sum := nums[first] + nums[left] + nums[right]
			if sum == 0 {
				for nums[left] == nums[left+1] {
					left++
				}
				for nums[right] == nums[right-1] {
					right--
				}
				res = append(res, []int{nums[first], nums[left], nums[right]})
				left++
				right--
			} else if sum > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return res
}

func ThreeSum2(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	// 排序
	sort.Ints(nums)
	ans := make([][]int, 0)
	for first := 0; first < len(nums); first++ {
		if nums[first] > 0 {
			break
		}
		left := first + 1
		right := len(nums) - 1

		for left < right {
			sum := nums[first] + nums[left] + nums[right]
			if sum == 0 {
				// 符合条件
				for nums[left+1] == nums[left] {
					// 去重
					left++
				}
				for nums[right-1] == nums[right] {
					// 去重
					right--
				}
				ans = append(ans, []int{nums[first], nums[left], nums[right]})
				left++
				right--
			} else if sum > 0 {
				right--
			} else {
				left++
			}
		}
	}
	return ans
}
