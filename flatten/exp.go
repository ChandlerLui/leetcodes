package flatten

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	head := &TreeNode{}
	var inOrder func(node *TreeNode)
	inOrder = func(node *TreeNode) {
		if node == nil {
			return
		}
		inOrder(node.Left)
		head.Right = node
		head.Left = nil
		head = node
		inOrder(node.Right)
	}
	inOrder(root)
}
