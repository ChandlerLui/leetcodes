package ImplenmentTrie

type Trie struct {
	M     map[byte]*Trie
	IsEnd bool
}

func Constructor() Trie {
	return Trie{
		M: make(map[byte]*Trie),
	}
}

func (this *Trie) Insert(word string) {
	//
	current := this
	for i := 0; i <= len(word); i++ {
		w := word[i]
		if v, ok := current.M[w]; ok {
			// 存在则进一步
			current = v
		} else {
			var isEnd bool
			if i == len(word)-1 {
				isEnd = true
			}
			// 不存在 则声明，然后进一步
			current.M[w] = &Trie{
				M:     make(map[byte]*Trie),
				IsEnd: isEnd,
			}
			current = current.M[w]
		}
	}
}

func (this *Trie) Search(word string) bool {
	for i := 0; i <= len(word); i++ {

	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	return false
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
