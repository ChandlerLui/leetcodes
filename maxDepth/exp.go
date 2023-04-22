package maxDepth

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	// 广度优先遍历
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	var ans int

	for len(queue) > 0 {
		size := len(queue)
		for size > 0 {
			// 弹出
			item := queue[0]
			queue = queue[1:]
			if item.Left != nil {
				queue = append(queue, item.Left)
			}
			if item.Right != nil {
				queue = append(queue, item.Right)
			}
			size--
		}
		ans++
	}
	return ans
}
