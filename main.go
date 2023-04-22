package main

import (
	"awesomeProject/LinkListPartition"
	"awesomeProject/MaxNoRepeatSubString"
	"fmt"
	"math"
	"unicode"
)

func main() {
	// l1 := []int{1,0,0,0,0,0,0,0,0,0,0,0,9,9,9,9,9,9,9}
	// l2 := []int{9,9,9,9}
	//var a  *addTwoNumbers.ListNode
	//for i := len(l1)-1; i>=0;i--{
	//	a = &addTwoNumbers.ListNode{
	//		Val:  l1[i],
	//		Next: a,
	//	}
	//}
	//var b  *addTwoNumbers.ListNode
	//for i := len(l2)-1; i>=0;i--{
	//	b = &addTwoNumbers.ListNode{
	//		Val:  l2[i],
	//		Next: b,
	//	}
	//}
	//
	//res := addTwoNumbers.AddTwoNumbers(a,b)
	//for {
	//	fmt.Println(res.Val)
	//	if res.Next == nil{
	//		break
	//	}
	//	res = res.Next
	//}

	fmt.Println(MaxNoRepeatSubString.LengthOfLongestSubstring1("abcba"))
	//a := []int{5,4}
	//var pre *reverseLinkedList.ListNode
	//for _, i := range a{
	//	a := &reverseLinkedList.ListNode{
	//		Val:  i,
	//		Next: pre,
	//	}
	//	pre = a
	//}
	//reverseLinkedList.ReverseBetween(pre, 1,2)
	//res := PalindromeNumber.IsPalindrome(11)
	//fmt.Println(res)

	//fmt.Println(ReverseInt.Reverse(121212))
	//fmt.Println(StringToInt.MyAtoi("9223372036854775808"))

	//fmt.Println(MaxContainer.MaxArea([]int{1,8,6,2,5,4,8,3,7}))
	//fmt.Println(LongestCommonPrefix.LongestCommonPrefix([]string{"dog","racecar","car"}))
	//ReplaceWords.ReplaceWords([]string{"cat","bat","rat"}, "the cattle was rattled by the battery")
	//for i:=0;i<=1000;i++{
	//	fmt.Println(i)
	//}
	var (
		pre *LinkListPartition.ListNode
		//b *hasCycle.ListNode
	)

	for _, item := range []int{5, 4, 3, 2, 1} {
		cur := &LinkListPartition.ListNode{Next: pre, Val: item}
		pre = cur
	}

	//reverseLinkedList.ReverseLinkListII(pre, 2,4)
	//RemoveNthLinkedList.RemoveNthFromEnd(pre, 2)
	//SortedListToBST.SortedListToBST(pre)
	//c := hasCycle.HasCycle(pre)
	//fmt.Println(c)
	//SearchAndInsert.SearchInsert([]int{1,2,4,5}, 3)
	//l := LRUCache.Constructor(3)
	//a := l.Get(1)
	//fmt.Println(a)
	//l.Put(2,3)
	//b:= l.Get(2)
	//fmt.Println(b)
	//l.Put(2,3)
	//l.Put(3,3)
	//l.Put(4,3)
	//l.Put(5,3)
	//fmt.Println(l)

	//fmt.Println(findDuplicate.FindDuplicate([]int{1,4,3,3,1}))
	// "{"1"：5，"2"：0}"
	//numTrees.NumTrees(5)
	//MaxSubArray.MaxSubArray([]int{1,2,-4,4,5})
	//LengthOfLIs.LengthOfLIS([]int{10,9,2,5,3,7,101,18})
	//LongestCommonSubSequence.LongestCommonSubsequence("abcde","ace")

	//LinkListPartition.Partition(pre, 3)
	//threeSum.ThreeSum([]int{-1,0,1,2,-1,-4})
	//sortColors.SortColors([]int{2,1,2,1,0,0})
	//a := removeElement.RemoveElement([]int{1,1,2,2,3,3}, 3)
	//b := removeElement.RemoveElement2([]int{1,1,2,2,3,3}, 3)
	//fmt.Println(b)
	//fmt.Println(a)
	////fmt.Println(fmt.Sprintf("%05d", 1))

	CheckDmMsg("1233")
	////math.MaxInt
	////conbineSUm.CombinationSum([]int{2,2,3,7},7)
	//fmt.Println( a1())
	//Permute.Permute([]int{1,2,3})
	//combine.Combine(4,2)

	//combinationSum3.CombinationSum3(3,7)
	//subSets.Subsets([]int{1,2,3})
	//searchRange.SearchRange([]int{1}, 1)
	//Search.Search([]int{4,1}, 1)
	//trap.Trap2([]int{0,1,0,2,1,0,1,3,2,1,2,1})

	//res := similarity([]int{1,2,3,4,5,6,7,8,9,10,11,23,33,44},[]int{1,2,3,4,5,6,7,9,8,10,11,23,33,44})
	//fmt.Println(res)
	a := []int{1, 2, 3}
	// [)
	fmt.Println(a[:1])
}

type Conf struct {
	EnableAll  bool           `json:"enable_all"`  // 是否全量
	WhiteUsers []int64        `json:"white_users"` // 白名单用户，白名单默认切全量
	AreaIds    []int64        `json:"area_ids"`    // 灰度期间的areaID
	AbEnds     map[int64]bool `json:"ab_ends"`
}

func Add(c *Conf) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	// here I get the attr from a nil struct
	fmt.Println(c.EnableAll)
}

func InArrayInt64(need int64, needArr []int64) bool {
	for _, v := range needArr {
		if need == v {
			return true
		}
	}
	return false
}

func similarity(slice1, slice2 []int) float64 {
	n, m := len(slice1), len(slice2)
	if n == 0 && m == 0 {
		return 1
	}
	if n == 0 || m == 0 {
		return 0
	}
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		dp[i][0] = i
	}
	for j := 1; j <= m; j++ {
		dp[0][j] = j
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if slice1[i-1] == slice2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1])) + 1
			}
		}
	}
	distance := dp[n][m]
	return math.Round((1 - float64(distance)/float64(max(n, m))) * 100)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func a1() int {
	a := 1
	defer func() {
		a = 2
	}()
	return a
}

func CheckDmMsg(msg string) (isValid bool) {
	var (
		cnt   int
		isHan bool
	)
	for _, s := range msg {
		if cnt >= 3 && isHan {
			isValid = true
			break
		}
		cnt++
		if unicode.Is(unicode.Han, s) {
			isHan = true
			continue
		}
		if unicode.IsDigit(s) || unicode.IsLetter(s) || unicode.IsPunct(s) {
			continue
		}
	}
	return
}
