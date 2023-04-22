package ReplaceWords

import (
	"strings"
)

//type Trie struct {
//	children [26]*Trie
//	word     string
//	isEnd    bool
//}
//
//// 648. Replace Words
//// 648. 单词替换
//// 思路：trie
//// time O(n) space O(n)
//func ReplaceWords(dictionary []string, sentence string) string {
//	// put all the roots in a trie
//	root := &Trie{}
//	for _, word := range dictionary {
//		cur := root
//		for _, ch := range word {
//			if cur.children[ch-'a'] == nil {
//				cur.children[ch-'a'] = &Trie{}
//			}
//			cur = cur.children[ch-'a']
//		}
//		cur.word = word
//		cur.isEnd = true
//	}
//
//	// iterate trie and match the prefix of a word
//	answer := &strings.Builder{}
//	s := strings.Split(sentence, " ")
//	for _, word := range s {
//		if len(answer.String()) > 0 {
//			answer.WriteString(" ")
//		}
//		cur := root
//		for _, ch := range word {
//			// 不匹配，表示当前单词不用替换，直接拼接即可。
//			if cur.children[ch-'a'] == nil {
//				break
//			}
//			// 匹配一个词根，则拼接词跟即可。
//			if cur.isEnd {
//				break
//			}
//			cur = cur.children[ch-'a']
//		}
//		if cur.isEnd {
//			answer.WriteString(cur.word)
//		} else {
//			answer.WriteString(word)
//		}
//	}
//	return answer.String()
//}

//func ReplaceWords(dict []string, sentence string) string {
//	trie := Constructor()
//	for _, v := range dict {
//		trie.Insert(v)
//	}
//	strs := split(sentence)
//	for i := 0; i < len(strs); i++ {
//		if s := trie.Search(strs[i]); s != "" {
//			strs[i] = s
//		}
//	}
//	return join(strs)
//}
//
//type Trie struct {
//	m map[byte]*Trie
//}
//
//func Constructor() Trie {
//	return Trie {
//		m: make(map[byte]*Trie),
//	}
//}
//
//func (this *Trie) Insert(word string)  {
//	current := this
//	for i := 0; i < len(word); i++ {
//		w := word[i]
//		if v, ok := current.m[w]; ok {
//			current = v
//		} else {
//			current.m[w] = &Trie{m: make(map[byte]*Trie)}
//			current = current.m[w]
//		}
//	}
//	// represents the end of the string
//	current.m['0'] = &Trie{m: make(map[byte]*Trie)}
//}
//
//func (this *Trie) Search(word string) string {
//	current, res := this, ""
//	for i := 0; i < len(word); i++ {
//		v, ok := current.m[word[i]]
//		if !ok {
//			return ""
//		}
//
//		res = res + string(word[i])
//		// return result if the end of the string
//		if v.m['0'] != nil {
//			return res
//		}
//		current = v
//	}
//	return res
//}
//
//func split(str string) []string {
//	res, i, s := []string{}, 0, 0
//	for i < len(str) {
//		if str[i] == ' ' {
//			res = append(res, str[s:i])
//			s = i+1
//		}
//		i++
//	}
//	return append(res, str[s:i])
//}
//
//func join(s []string) string {
//	res := ""
//	for _, v := range s {
//		if res == "" {
//			res = v
//		} else {
//			res = res + " " + v
//		}
//	}
//	return res
//}

type Trie struct {
	M     map[byte]*Trie
	IsEnd bool
}

// ReplaceWords trie数 实现单词替换
func ReplaceWords(dictionary []string, sentence string) string {
	// step1 遍历dictionary生成字典树
	// step2 遍历sentence ，匹配前段的数据，命中返回
	// step3 拼接字符串
	trie := &Trie{M: map[byte]*Trie{}}
	// 遍历字典
	for _, item := range dictionary {
		// insert Trie
		current := trie
		for i := 0; i < len(item); i++ {
			word := item[i]
			if v, ok := current.M[word]; ok {
				// 如果trie树中存在这个元素，则进一步
				current = v
			} else {
				var isEnd bool
				if i == len(item)-1 {
					isEnd = true
				}
				current.M[word] = &Trie{M: map[byte]*Trie{}, IsEnd: isEnd}
				current = current.M[word]
			}
		}
	}

	splits := strings.Split(sentence, " ")
	resList := make([]string, 0, len(splits))
	for _, words := range splits {
		var res string
		current := trie
		for i := 0; i < len(words); i++ {
			word := words[i]

			v, ok := current.M[word]
			if !ok {
				res = ""
				break
			}
			if current.IsEnd {
				break
			}
			res += string(word)
			current = v
		}
		if res == "" {
			res = words
		}
		resList = append(resList, res)
	}
	return strings.Join(resList, " ")
}
