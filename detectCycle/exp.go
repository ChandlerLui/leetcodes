package detectCycle

import "awesomeProject/hasCycle"

func detectCycle(head *hasCycle.ListNode) *hasCycle.ListNode {
	// 方法1：hash表
	existMap := make(map[*hasCycle.ListNode]struct{})
	for head != nil {
		if _, exist := existMap[head]; exist {
			return head
		}
		existMap[head] = struct{}{}
		head = head.Next
	}
	return nil
}

func detectCycle2(head *hasCycle.ListNode) *hasCycle.ListNode {
	// 方法2:快慢指针
	// 环形链表1的快慢指针重合后，快指针走的步数是慢指针的2倍。 fast := 2*slow
	// fast 和 slow重合，除去环前的节点，fast走过的步数比slow多走了N环 及 fast ：= slow + N * cycle
	// so slow = N * cycle
	// slow 走到环的第一个节点的需要 a + N * Cycle
	if head == nil {
		return head
	}

	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			// 一定有环， head 再走N步
			for {
				if head == slow {
					return head
				}
				head = head.Next
				slow = slow.Next
			}
		}
	}
	return nil
}
