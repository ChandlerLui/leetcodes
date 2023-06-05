package reverseGroup

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseKGroup(head *ListNode, k int) *ListNode {
	var length int
	temp := head

	for temp != nil {
		length++
		temp = temp.Next
	}

	dummyNode := &ListNode{Next: head}
	p0 := dummyNode

	var pre *ListNode
	cur := p0.Next
	//大于k遍历
	for length >= k {
		length -= k
		for i := 0; i < k; i++ {
			//走k步
			next := cur.Next
			cur.Next = pre
			pre = cur
			cur = next
		}

		n := p0.Next
		p0.Next.Next = cur
		p0.Next = pre
		p0 = n
	}
	return dummyNode.Next
}
