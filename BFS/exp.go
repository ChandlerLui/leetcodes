package BFS

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 广度优先遍历
//func levelOrder(root *TreeNode) [][]int {
//	if root == nil {
//		return nil
//	}
//	allRes := make([][]int,0)
//	queue := []*TreeNode{root}
//	for i:=0;len(queue)>0;i++{
//		// i 表示第几层，遍历终止条件是 队列为空
//		res := make([]int, 0)// 每一层的结果
//		NextQueue := make([]*TreeNode, 0)
//		for j:=0;j<len(queue);j++{
//			// 从里面取出节点
//			node := queue[j]
//			res = append(res, node.Val)
//			if node.Left != nil{
//				NextQueue = append(NextQueue, node.Left)
//			}
//			if node.Right != nil{
//				NextQueue = append(NextQueue, node.Right)
//			}
//		}
//		queue = NextQueue
//	}
//	return allRes
//}
//
//
// 广度优先
