package addTwoNumbers

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
* Definition for singly-linked list.
 type ListNode struct {
    Val int
     Next *ListNode
}
*/
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}
	var (
		l1Str string
		l2Str string
	)
	//
	for {
		l1Str = strconv.FormatInt(int64(l1.Val), 10) + l1Str
		if l1.Next == nil {
			break
		}
		l1 = l1.Next
	}

	for {
		l2Str = strconv.FormatInt(int64(l2.Val), 10) + l2Str
		if l2.Next == nil {
			break
		}
		l2 = l2.Next
	}
	var (
		res      *ListNode
		previous *ListNode
	)
	l2Int, err := strconv.ParseUint(l2Str, 10, 64)
	if err != nil {
		return nil
	}
	l1Int, err := strconv.ParseUint(l1Str, 10, 64)
	if err != nil {
		return nil
	}
	for _, i := range strconv.FormatUint(l2Int+l1Int, 10) {
		a, err := strconv.ParseInt(fmt.Sprintf("%c", i), 10, 64)
		if err != nil {
			return nil
		}
		res = &ListNode{
			Val:  int(a),
			Next: previous,
		}
		previous = res
	}
	return res
}
