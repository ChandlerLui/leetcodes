package buildTree

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	// 从 前序遍历 和 中序遍历 构建二叉树
	// preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
	// 输出: [3,9,20,null,null,15,7]
	// 前序: root -> left -> right
	// 中序: left -> root -> right
	if len(preorder) == 0 {
		return nil
	}
	i := 0
	for ; i < len(preorder); i++ {
		// 通过preorder的0号位
		if preorder[0] == inorder[i] {
			break
		}
	}
	//
	root := &TreeNode{Val: preorder[0]}
	root.Left = buildTree(preorder[1:len(inorder[:i])+1], inorder[:i])
	root.Right = buildTree(preorder[len(inorder[:i])+1:], inorder[i+1:])
	return root
}
