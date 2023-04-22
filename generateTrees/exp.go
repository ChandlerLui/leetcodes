package generateTrees

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func generateTrees(n int) []*TreeNode {
	// 给定某个节点，生存
	// 暴力回shu
	if n == 0 {
		return nil
	}
	//
	return helper(1, n)

}

func helper(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{}
	}
	allTrees := make([]*TreeNode, 0)
	for i := start + 1; i <= end; i++ {
		leftTrees := helper(start, i-1)
		RightTrees := helper(i+1, end)
		for _, left := range leftTrees {
			for _, right := range RightTrees {
				t := &TreeNode{Val: i}
				t.Left = left
				t.Right = right
				allTrees = append(allTrees, t)
			}
		}
	}
	return allTrees
}
