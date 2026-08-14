type Node struct {
	next *Node
	prev *Node
	key  int
	val  int
}

type LRUCache struct {
	start *Node
	end   *Node
	nums  map[int]*Node
	cap   int
}

func Constructor(capasity int) LRUCache {
	start := &Node{}
	end := &Node{}
	start.next = end
	end.prev = start
	return LRUCache{
		start: start,
		end:   end,
		cap:   capasity,
		nums:  make(map[int]*Node, capasity),
	}
}

func (c *LRUCache) insert(key, value int) {
	node := &Node{key: key, val: value}
	end := c.end
	newPrev := end.prev

	node.next = end
	node.prev = newPrev
	end.prev = node
	newPrev.next = node

	c.nums[key] = node
}

func (c *LRUCache) delete(key int) {
	node := c.nums[key]
	prev := node.prev
	next := node.next
	prev.next = next
	next.prev = prev
	delete(c.nums, key)
}

func (c *LRUCache) Get(key int) int {
	node, ok := c.nums[key]
	if ok {
		c.delete(key)
		c.insert(key, node.val)
		return node.val
	}
	return -1
}

func (c *LRUCache) Put(key, value int) {
	hasNode, ok := c.nums[key]
	if ok {
		c.delete(hasNode.key)
	}
	c.insert(key, value)
	if len(c.nums) > c.cap {
		c.delete(c.start.next.key)
	}
}
