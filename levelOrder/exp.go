package levelOrder

import "container/list"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	var ans [][]int
	queue := list.New()
	queue.PushBack(root)

	for queue.Len() > 0 {
		var res []int
		size := queue.Len()

		for i := 0; i < size; i++ {
			node := queue.Remove(queue.Front()).(*TreeNode)
			res = append(res, node.Val)

			if node.Left != nil {
				queue.PushBack(node.Left)
			}
			if node.Right != nil {
				queue.PushBack(node.Right)
			}
		}
		ans = append(ans, res)
	}

	return ans
}
