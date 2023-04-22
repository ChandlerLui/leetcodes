package isValidTree

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
   5
4    7
   6   8
*/

// stack：[6] root：3
//
var pre = math.MinInt64

func isValidBST(root *TreeNode) bool {
	// 中序遍历
	if root == nil {
		return true
	}
	// zuo
	if !isValidBST(root.Left) {
		return false
	}
	if root.Val <= pre {
		return false
	}
	pre = root.Val
	return isValidBST(root.Right)
}
