package quicksort

import "math/rand"

func QuickSortIterative(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	ans := make([]int, 0, len(arr))
	stack := [][]int{arr}

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if len(curr) <= 1 {
			if len(curr) == 1 {
				ans = append(ans, curr[0])
			}
			continue
		}

		pivotIndex := rand.Intn(len(curr))
		pivot := curr[pivotIndex]
		curr[pivotIndex], curr[0] = curr[0], curr[pivotIndex]
		var left, right []int

		for _, v := range curr[1:] {
			if v < pivot {
				left = append(left, v)
			} else {
				right = append(right, v)
			}
		}

		left = append(left, pivot)
		stack = append(stack, right, left)
	}

	return ans
}

// iter way
func IterQuickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	ans := make([]int, 0, len(nums))
	// 模拟递归栈
	stack := [][]int{nums}

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if len(curr) <= 1 {
			if len(curr) == 1 {
				ans = append(ans, curr[0])
			}
		}

		rd := rand.Intn(len(curr) - 1)
		pivot := curr[rd]
		curr[rd], curr[0] = curr[0], curr[rd]

		left, right := []int{}, []int{}
		for _, item := range curr[1:] {
			if item < pivot {
				left = append(left, item)
			} else {
				right = append(right, item)
			}
		}
		left = append(left, pivot)
		stack = append(stack, right, left)

	}
	return ans
}

func RecursionQuickSort(nums []int) []int {
	if len(nums) <= 1 {
		return nums
	}
	ans := make([]int, 0, len(nums))
	rd := rand.Intn(len(nums) - 1)
	pivot := nums[rd]
	nums[pivot], nums[0] = nums[0], nums[pivot]

	left, right := []int{}, []int{}
	for _, item := range nums[1:] {
		if item < pivot {
			left = append(left, item)
		} else {
			right = append(right, item)
		}
	}
	left = append(left, pivot)

	ans = append(ans, RecursionQuickSort(left)...)
	ans = append(ans, RecursionQuickSort(right)...)
	return ans
}
