package other

type Node struct {
	Key int
	Value int
	Prev *Node
	Next *Node
}


type LRUCache struct {
	Size int
	Capacity int
	Cache map[int]*Node
	Head, Tail *Node
}

func Constructor ()*LRUCache{
	head, tail := &Node{}, &Node{}
	head.Next, tail.Prev = tail,head
	return &LRUCache{
		Size: 0,
		Capacity: 10,
		Cache: make(map[int]*Node),
		Head: head,
		Tail: tail,
	}
}

func (l *LRUCache)Get(key int)int{
	if node ,ok := l.Cache[key]; ok{
		// 把该节点移到双向链表的前端
		l.MoveToHead(node)
		return node.Value
	}
	return -1
}
func (l *LRUCache)Put(key int, val int){
	if node,ok:= l.Cache[key];ok{
		node.Value = val
		l.MoveToHead(node)
	}
	node := &Node{
		Key:   key,
		Value: val,
		Prev:  nil,
		Next:  nil,
	}
	l.Size++
	l.Cache[key]=node
	l.AddToHead(node)
	if l.Size > l.Capacity{
		removed := l.RemoveTail()
		l.RemoveNode(l.Tail)
	}
}

func (l *LRUCache) MoveToHead(node *Node){
	l.RemoveNode(node)
	l.AddToHead(node)
}

func (l *LRUCache)RemoveNode(node *Node){
	node.Prev.Next = node.Next
	node.Next.Prev= node.Prev
}

func (l *LRUCache)AddToHead(node *Node){
	node.Next= l.Head
	l.Head.Prev = node
	l.Tail.Prev = node
	l.Head = node
}
func (l *LRUCache)RemoveTail()*Node{
	node := l.Tail.Prev
}