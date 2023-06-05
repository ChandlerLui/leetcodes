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

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	dummyNode := &ListNode{Next: head}
	p0 := dummyNode
	for i := 0; i < left; i++ {
		// 走left-1步
		p0 = p0.Next
	}
	// p0是left的前一个，p0。next表示 反转连标的第一个节点
	// 执行反转
	var pre *ListNode
	cur := p0.Next // 指向反转连标的第一个节点
	for i := 0; i < right-left+1; i++ {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	p0.Next.Next = cur
	p0.Next = pre
	return dummyNode.Next
}
