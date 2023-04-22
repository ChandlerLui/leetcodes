package kthSmallest

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func kthSmallest(root *TreeNode, k int) int {
	// 中序遍历
	ans := make([]int, 0)
	var inOrder func(node *TreeNode)

	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inOrder(node.Left)
		ans = append(ans, node.Val)
		inOrder(node.Right)
	}

	inOrder(root)
	if k > len(ans) {
		return -1
	}
	return ans[k-1]
}
