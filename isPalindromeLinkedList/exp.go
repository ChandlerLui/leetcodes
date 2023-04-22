package isPalindromeLinkedList

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	if head.Next.Next == nil {
		return head.Val == head.Next.Val
	}
	// 判断回文链表
	st := []*ListNode{head}
	fast, slow := head, head
	for fast.Next != nil || fast.Next.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		st = append(st, slow)
	}
	if fast.Next == nil {
		// 总长度是奇数，mid是中间唯一
		i := len(st) - 2
		for slow.Next != nil {
			if slow.Next.Val != st[i].Val || i < 0 {
				return false
			}
			slow = slow.Next
			i--
		}
	} else {
		i := len(st) - 1
		// 偶数 长度
		for slow.Next != nil {
			if slow.Next.Val != st[i].Val || i < 0 {
				return false
			}
			slow = slow.Next
			i--
		}
	}
	return true
}
