package hasCycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func HasCycle(head *ListNode) bool {
	// 快慢指针法
	// 1：申明一个快慢指针
	dummyNode := &ListNode{Next: head}
	slow, fast := dummyNode, dummyNode
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			return true
		}
	}
	return false
}
