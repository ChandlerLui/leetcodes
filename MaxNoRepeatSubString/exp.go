package MaxNoRepeatSubString

func LengthOfLongestSubstring(s string) string {
	length := len(s)
	if length <= 1 {
		return s
	}
	res := make([][]bool, 0)
	// 使用二维数据记录 s[i,j],是否是回文
	for i := 0; i < length; i++ {
		inner := make([]bool, len(s))
		res = append(res, inner)
	}
	// 将每个字符置成true
	for i := 0; i < length; i++ {
		res[i][i] = true
	}
	var (
		maxLen = 1
		begin  = 0
	)
	// 列举回文子串的长度,因为子串长度小的是回文串,包含他的才可能是回文串
	for l := 2; l <= length; l++ {
		// 枚举所有子串
		for left := 0; left < length; left++ {
			// left 是 左边界, 右边界 是 子串长度 + 左边界
			right := left + l - 1

			if right >= length {
				break
			}
			if s[left] == s[right] {
				if l > 2 {
					res[left][right] = res[left+1][right-1]
				} else {
					res[left][right] = true
				}
			}

			if l > maxLen && res[left][right] {
				maxLen = l
				begin = left
			}

		}
	}
	return s[begin : begin+maxLen]
}

//LengthOfLongestSubstring1 最长回文子串
func LengthOfLongestSubstring1(s string) string {

	// 动态规划5步：
	// 1：确定dp数组以及对应下标的含义
	// 2：确定状态转换公式
	// 3：dp数组如何初始化（初始化需要注意）
	// 4：for循环的遍历顺序，初始值，终止条件等。
	// 5：打印dp数组，确定逻辑正确。

	// 我们使用dp[i][j]二维数组表示，i，j的子串为回文子串，用true代替
	// i j 表示字符串的左右下标
	// 若dp[i][j] 为回文子串，则dp[i-1][j+1]也是回文子串
	// 所以dp[i][j] = dp[i-1][j+1] (i>1)
	// 如果
	// a b c b a
	if len(s) < 2 {
		return s
	}
	dpList := make([][]bool, len(s))
	for index := range dpList {
		dpList[index] = make([]bool, len(s))
		dpList[index][index] = true
	}
	var (
		max, begin int
	)
	// 对长度进行枚举
	for length := 2; length <= len(s); length++ {
		for left := 0; left < len(s); left++ {
			right := left + length - 1
			if right >= len(s) {
				break
			}
			if right < left {
				break
			}
			if s[left] == s[right] {
				if length == 2 {
					dpList[left][right] = true
				} else {
					dpList[left][right] = dpList[left+1][right-1]
				}
			}
			if length > max && dpList[left][right] {
				max = length
				begin = left
			}
		}
	}
	return s[begin : begin+max]
}
