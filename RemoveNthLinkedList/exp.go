package RemoveNthLinkedList

import "awesomeProject/reverseLinkedList"

// 移除链表的第N个元素，并返回链表的头节点,快慢指针法

//func RemoveNthFromEnd(head *reverseLinkedList.ListNode, n int) *reverseLinkedList.ListNode {
//	dummy := &reverseLinkedList.ListNode{Next: head}
//	left, right := dummy, dummy
//	for ; n > 0; n-- {
//		right = right.Next
//	}
//	for right.Next != nil {
//		left = left.Next
//		right = right.Next
//	}
//	left.Next = left.Next.Next
//	return dummy.Next
//}

func RemoveNthFromEnd(head *reverseLinkedList.ListNode, n int) *reverseLinkedList.ListNode {
	// 移除倒数第n个链表
	var dummyNode = &reverseLinkedList.ListNode{Next: head}
	left, right := dummyNode, dummyNode
	// 先将快指针 right往后移动n个位置，
	for i := 0; i < n; i++ {
		right = right.Next
	}
	// 慢指针开始移动
	for right.Next != nil {
		left = left.Next
		right = right.Next
	}
	left.Next = left.Next.Next
	return dummyNode.Next
}
