package main

import (
	"awesomeProject/LinkListPartition"
	"awesomeProject/MaxNoRepeatSubString"
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
	"sync"
	"unicode"
)

func merge(nums [][]int) {
	//
	fmt.Println(nums)
	ans := make([]int, 0)
	for len(nums) > 1 {
		//	按最后一位排序
		sort.Slice(nums, func(i, j int) bool {
			if len(nums[i]) == 0 {
				return false
			} else if len(nums[j]) == 0 {
				return true
			}
			return nums[i][len(nums[i])-1] > nums[j][len(nums[j])-1]
		})

		for index := 0; index < len(nums); index++ {
			if len(nums[index]) == 0 {
				nums = nums[:index]
				break
			}
			ans = append(ans, nums[index][len(nums[index])-1])
			nums[index] = nums[index][:len(nums[index])-1]
		}
	}
	// 最后一个
	ans = append(ans, nums[0]...)
	fmt.Println(ans)
}

func main() {

	// 有n个不等长升须的数组，合并的降序的数组
	input := make([][]int, 0)
	input = [][]int{{1, 3, 5}, {2, 4}, {6}, {}}

	// [1,3,5]
	// [2,4]
	// [6]
	merge(input)
	//a := "ALTER TABLE ap_user_title_%d ADD lock_status tinyint NOT NULL DEFAULT 0 COMMENT '锁定状态 1 锁定 0 未锁定';"
	//for i:=0;i<=99;i++{
	//	fmt.Println(fmt.Sprintf(a, i))
	//}
	//m := make(map[int]int, 10)
	//for i := 1; i<= 10; i++ {
	//	m[i] = i
	//}
	//
	//for k, v := range(m) {
	//	go func() {
	//		fmt.Println("k ->", k, "v ->", v)
	//	}()
	//}
	//removeDuplicates.IsIsomorphic("egg",  "ada")
	fmt.Println(1 << 3)

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

	// [)

}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func extractNumbers(s string) []int {
	re := regexp.MustCompile(`\d+`)
	matches := re.FindAllString(s, -1)
	nums := make([]int, len(matches))

	for i, match := range matches {
		num, _ := strconv.Atoi(match)
		nums[i] = num
	}

	return nums
}

func Helper() {

	limit := make(chan struct{}, 30)

	wg := sync.WaitGroup{}
	for i := 0; i < 99; i++ {
		limit <- struct{}{}
		go func(i int) {
			wg.Add(1)
			defer wg.Done()

			fmt.Println(i)
			<-limit
		}(i)
	}
}

type TitleItem struct {
	// ap_title字段
	Id                int64  `json:"id"`
	TitleId           int64  `json:"title_id"`
	Level             int64  `json:"level"`
	NormalTag         int64  `json:"normal_tag"`
	SpecialTag        string `json:"special_tag"`
	Name              string `json:"name"`
	Url               string `json:"url"`
	H5Url             string `json:"h5_url"`
	MobilePicUrl      string `json:"mobile_pic_url"`
	WebPicUrl         string `json:"web_pic_url"`
	Source            string `json:"source"`
	Description       string `json:"description"`
	OldIdentification string `json:"old_identification"`
	Status            int    `json:"status"`
	HWidth            int64  `json:"h_width"`
	HHeight           int64  `json:"h_height"`
	XhWidth           int64  `json:"xh_width"`
	XhHeight          int64  `json:"xh_height"`
	XxhWidth          int64  `json:"xxh_width"`
	XxhHeight         int64  `json:"xxh_height"`
	Colorful          int64  `json:"colorful"`
	Grade             int64  `json:"grade"`
	Effective         int64  `json:"effective"`
	// cid不是ap_title字段，根据头衔的子头衔等级设置的值
	Cid int64 `json:"cid"`
	// 子头衔，根据原始数据level等逻辑拼装而成，部分头衔才有子头衔
	SubTitle []*TitleItem `json:"sub_title"`
	// 同一个装扮，内部等级
	TitleLevel int32 `json:"title_level"`
	// 装扮获得时，是否主动提示用户配置, 默认不提示
	IsNeedNotify int32 `json:"is_need_notify"`
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
