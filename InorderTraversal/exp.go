package InorderTraversal

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func inorderTraversal(root *TreeNode) []int {
	// 中序遍历是由左-》根-》右，所以添加的顺序
	// 递归终止条件是node= nil
	//
	return inOrder(root)
}

func inOrder(node *TreeNode) (res []int) {
	res = make([]int, 0)
	if node == nil {
		return
	}
	// 递归左节点
	res = append(res, inOrder(node.Left)...)
	res = append(res, node.Val)
	res = append(res, inOrder(node.Right)...)
	return
}

/*
class Solution:
    def inorderTraversal(self, root: TreeNode) -> List[int]:
        WHITE, GRAY = 0, 1
        res = []
        stack = [(WHITE, root)]
        while stack:
            color, node = stack.pop()
            if node is None: continue
            if color == WHITE:
                stack.append((WHITE, node.right))
                stack.append((GRAY, node))
                stack.append((WHITE, node.left))
            else:
                res.append(node.val)
        return res
*/

type ColorNode struct {
	Color int
	*TreeNode
}

//func inorderTraversal2(root *TreeNode) []int {
//	// 颜色标记法
//	// 已经遍历过的节点标记为灰色，新节点标记为白色
//	//
//	// 申明一个栈
//	stack := []*ColorNode{{
//		Color:    0,
//		TreeNode: root,
//	},}
//	for ;len(stack)!= 0;{
//
//	}
//}
