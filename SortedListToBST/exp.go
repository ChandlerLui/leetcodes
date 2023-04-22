package SortedListToBST

type ListNode struct {
	Val  int
	Next *ListNode
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 将有序链表转换成二叉搜索树
func SortedListToBST(head *ListNode) {
	// 1:快慢指针找到中间节点
	// 2：将链表分为 left：中间，中间：right
	// 3：递归
	tree := buildTree(head, nil)
	print(tree)
}

func getMiddle(left, right *ListNode) *ListNode {
	slow, fast := left, left
	for fast != right && fast.Next != right { // 防止进两格，导致nil
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func buildTree(left, right *ListNode) *TreeNode {
	if left == right {
		return nil
	}
	mid := getMiddle(left, right)

	root := &TreeNode{
		Val: mid.Val,
	}
	root.Left = buildTree(left, mid)
	root.Right = buildTree(mid.Next, right)
	return root
}
