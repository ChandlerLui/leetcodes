package LRUCache

type LinkNode struct {
	Key, Value int
	Prev, Next *LinkNode
}

type LRUCache struct {
	Size, Capacity int
	Head, Tail     *LinkNode
	ExistMap       map[int]*LinkNode
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		Size:     0,
		Capacity: capacity,
		Head: &LinkNode{
			Key:   0,
			Value: 0,
			Prev:  nil,
			Next:  nil,
		},
		Tail: &LinkNode{
			Key:   0,
			Value: 0,
			Prev:  nil,
			Next:  nil,
		},
		ExistMap: map[int]*LinkNode{},
	}
	lru.Head.Next = lru.Tail
	lru.Tail.Prev = lru.Head
	return lru
}

func (this *LRUCache) Get(key int) int {
	// 先判断是否存在在lru缓存中
	// 存在直接返回。并且将对应的node提到head（需要移除后面的node）
	// else 直接返回
	if target, exist := this.ExistMap[key]; !exist {
		return -1
	} else {
		this.MoveToFront(target)
		return target.Value
	}
}

func (this *LRUCache) Put(key int, value int) {
	// 放入元素
	// 1：元素存在 则更新，并且提到前面
	// 2：元素不存在，则在头部加一个节点，加入map中
	//		2.1 判断是否超过长度，过长 则剔除尾节点
	if target, exist := this.ExistMap[key]; exist {
		target.Value = value
		this.MoveToFront(target)
	} else {
		node := &LinkNode{
			Key:   key,
			Value: value,
			Prev:  this.Head,
			Next:  this.Head.Next,
		}
		this.Head.Next.Prev = node
		this.Head.Next = node

		this.ExistMap[key] = node
		this.Size++

		if this.Size > this.Capacity {
			// 超额了，需要减掉
			lease := this.RemoveTail()
			delete(this.ExistMap, lease.Key)
		}

	}
}

// MoveToFront 节点插入头部
func (this *LRUCache) MoveToFront(node *LinkNode) {
	this.RemoveNode(node)
	node.Next = this.Head.Next
	node.Prev = this.Head
	this.Head.Next.Prev = node
	this.Head.Next = node

}

// RemoveNode 移除节点
func (this *LRUCache) RemoveNode(node *LinkNode) {
	// 移除双向链表中的某一个节点
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}

func (this *LRUCache) RemoveTail() (node *LinkNode) {
	// 移除尾节点
	node = this.Tail.Prev
	this.Tail.Prev.Prev.Next = this.Tail
	this.Tail.Prev = this.Tail.Prev.Prev
	return
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func permute(nums []int) (ans [][]int) {
	n := len(nums)
	path := make([]int, n)
	onPath := make([]bool, n)
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]int(nil), path...))
			return
		}
		for j, on := range onPath {
			if !on {
				path[i] = nums[j]
				onPath[j] = true
				dfs(i + 1)
				onPath[j] = false
			}
		}
	}
	dfs(0)
	return
}
