package lowestCommonAncestor

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// lowestCommonAncestor
// 使用hash表进行存储，
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	parentMap := make(map[int]*TreeNode)
	visited := make(map[int]bool)

	// 先进行递归遍历，存储所有树的parent节点
	// 找到p或q的前序节点
	var dfs func(node *TreeNode)

	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			parentMap[node.Left.Val] = node
			dfs(node.Left)
		}
		if node.Right != nil {
			parentMap[node.Right.Val] = node
			dfs(node.Right)
		}
	}
	dfs(root)
	for p != nil {
		visited[p.Val] = true
		p = parentMap[p.Val]
	}
	for q != nil {
		if visited[q.Val] {
			return q
		}
		q = parentMap[q.Val]
	}
	return nil
}
