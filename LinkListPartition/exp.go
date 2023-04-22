package LinkListPartition

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func Partition(head *ListNode, x int) *ListNode {
	if head == nil {
		return nil
	}
	// 1,2,3,4,5
	// front :-1 -> 1
	front := &ListNode{Val: -1}
	frontHead := front
	end := &ListNode{Val: -1}
	endHead := end
	for head != nil {
		if head.Val < x {
			front.Next = head
			front = front.Next
		} else {
			end.Next = head
			end = end.Next
		}
		head = head.Next
	}
	front.Next = endHead.Next
	end.Next = nil
	return frontHead.Next
}
