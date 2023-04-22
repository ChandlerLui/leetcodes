package reverseLinkedList

type ListNode struct {
	Val  int
	Next *ListNode
}

// ReverseBetween 反转单链表, 给一个链表, 反转[left, right]的链表
func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	// 特殊情况
	if head == nil || head.Next == nil || left > right {
		return head
	}
	// 首先 获取链表的left,right节点(注意记录left,right附近节点的记录问题
	// 声明 虚拟节点
	var dummyNode = &ListNode{
		Val: -1,
	}
	// 虚拟节点的下一节是头结点
	dummyNode.Next = head
	pre := dummyNode // 前置节点
	// 循环迭代到左节点的前一个节点
	for i := 1; i < left-1; i++ {
		pre = pre.Next
	}
	// 左节点是前置节点的下一个节点
	leftNode := pre.Next
	rightNode := leftNode // 从左节点开始迭代到右节点
	for i := 1; i < right-left+1; i++ {
		rightNode = rightNode.Next
	}

	tail := rightNode.Next // 记录右节点的下一个节点
	// 开始切断
	pre.Next = nil
	rightNode.Next = nil
	NewHead := Reverse(leftNode)
	pre.Next = NewHead
	leftNode.Next = tail

	return dummyNode.Next
}

// Reverse reverse List
func Reverse(head *ListNode) *ListNode {
	var pre *ListNode // 前置节点
	current := head   // 当前节点
	for current != nil {
		// 1 -> 2 -> 3 -> 4
		next := current.Next //2
		current.Next = pre
		pre = current
		current = next
	}
	return pre
}

// middle
// 0<-1-<2->3->4
// 保存1的下一节点，防止断链
// 声明当前节点的pre节点
// 获取
// 反转链表
//func ReverseList(head *ListNode) *ListNode {
//	if head == nil || head.Next == nil{
//		return head
//	}
//	var pre *ListNode // 前置节点
//	current := head // 1
//	for ;current != nil;{
//		next := current.Next // 2
//		current.Next = pre // 0
//		pre = current
//		current = next
//	}
//	return current
//}

func ReverseLinkListII(head *ListNode, left, right int) *ListNode {
	// 反转中间连标
	// 给一个left，right
	// 1：判断边界条件
	// 2：声明哨兵节点
	// 3：遍历找到左节点 及右节点，保存对应的场景
	// 4：切断，并执行反转，&& 拼接起来
	if head == nil || head.Next == nil || left >= right {
		return head
	}
	var dummyNode = &ListNode{
		Val: -1,
	}
	dummyNode.Next = head // 声明哨兵节点
	pre := dummyNode
	// 寻找左节点
	// dummy ->1->2
	for i := 1; i < left; i++ {
		pre = pre.Next
	}

	leftNode := pre.Next
	rightNode := leftNode

	// 寻找右节点
	// 1-2-3-4-5
	for i := 0; i < right-left; i++ {
		rightNode = rightNode.Next
	}

	tail := rightNode.Next

	// 切断
	pre.Next = nil
	rightNode.Next = nil

	pre.Next = R(leftNode)
	leftNode.Next = tail
	return dummyNode.Next
}

func R(head *ListNode) *ListNode {
	var pre *ListNode
	current := head
	for current != nil {
		next := current.Next
		current.Next = pre
		pre = current
		current = next
	}
	return pre
}
