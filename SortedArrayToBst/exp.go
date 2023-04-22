package SortedArrayToBst

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 有序列表转换为平衡🌲
func sortedArrayToBST(nums []int) *TreeNode {
	var trans func([]int, int, int) *TreeNode
	trans = func(nums []int, left int, right int) *TreeNode {
		if left > right {
			return nil
		}
		mid := (left + right) / 2
		root := &TreeNode{
			Val: nums[mid],
		}
		root.Left = trans(nums, left, mid-1)
		root.Right = trans(nums, mid+1, right)
		return root
	}
	return trans(nums, 0, len(nums)-1)
}
