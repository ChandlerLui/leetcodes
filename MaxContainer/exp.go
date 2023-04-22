package MaxContainer

//func MaxArea(height []int) int {
//	// step 1 conner case length = 0,1,2
//	// gen a  array array
//	// range
//	if len(height) == 0 || len(height) == 1{
//		return 0
//	}
//	// generate all steps data
//	var (
//		max,left,right int
//	)
//	right = len(height)-1
//	for left < right{
//		var r bool
//		lower := height[left]
//		if height[left] > height[right]{
//			lower = height[right]
//			r = true
//		}
//		if (right-left)*lower > max{
//			max = (right-left)*lower
//		}
//		if r{
//			right-=1
//		}else{
//			left +=1
//		}
//
//	}
//	return max
//}

func removeDuplicates(nums []int) int {
	//
	if len(nums) <= 1 {
		return len(nums)
	}
	// [1,1,2,2,3,3]
	fast, slow := 1, 1
	for ; fast < len(nums); fast++ {
		if nums[fast] != nums[fast-1] {
			nums[slow] = nums[fast]
			slow++
		}
	}
	return slow
}
