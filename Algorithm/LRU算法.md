

LRU算法的全称是“**Least Recently Used**”，即“最近最少使用”算法。LRU算法的基本思想是，当缓存空间已满时，优先淘汰最近最少使用的缓存数据，以腾出更多的缓存空间。因此，LRU算法需要维护一个缓存访问历史记录，以便确定哪些缓存数据是最近最少使用的。

LRU算法的实现可以采用多种数据结构，其中最常见的是使用一个**双向链表**和一个**哈希表**。**双向链表**用于维护缓存数据的访问顺序，**哈希表**用于快速查找缓存数据。具体来说，当新的数据被访问时，先在**哈希表**中查找该数据是否已经存在于缓存中，如果存在，则将该数据移动到**双向链表**的头部，表示该数据是最近访问的数据；如果不存在，则需要将该数据添加到缓存中，并将其添加到双向链表的头部。当缓存空间已满时，需要淘汰双向链表中最后一个节点，同时在哈希表中删除对应的缓存数据。

# Go语言实现
```go
package other  
  
type Node struct {  
   Key   int  
   Value int  
   Prev  *Node  
   Next  *Node  
}  
  
type LRUCache struct {  
   Size       int  
   Capacity   int  
   Cache      map[int]*Node  
   Head, Tail *Node  
}  
  
func Constructor() *LRUCache {  
   head, tail := &Node{}, &Node{}  
   head.Next, tail.Prev = tail, head  
   return &LRUCache{  
      Size:     0,  
      Capacity: 10,  
      Cache:    make(map[int]*Node),  
      Head:     head,  
      Tail:     tail,  
   }  
}  
  
func (l *LRUCache) Get(key int) int {  
   if node, ok := l.Cache[key]; ok {  
      // 把该节点移到双向链表的前端  
      l.MoveToHead(node)  
      return node.Value  
   }  
   return -1  
}  
func (l *LRUCache) Put(key int, val int) {  
   if node, ok := l.Cache[key]; ok {  
      node.Value = val  
      l.MoveToHead(node)  
      return  
   }  
   node := &Node{Key: key, Value: val}  
   l.Size++  
   l.Cache[key] = node  
   l.AddToHead(node)  
   if l.Size > l.Capacity {  
      removed := l.RemoveTail()  
      delete(l.Cache, removed.Key)  
      l.Size--  
   }  
}  
  
func (l *LRUCache) MoveToHead(node *Node) {  
   l.RemoveNode(node)  
   l.AddToHead(node)  
}  
  
func (l *LRUCache) RemoveNode(node *Node) {  
   node.Prev.Next = node.Next  
   node.Next.Prev = node.Prev  
}  
  
func (l *LRUCache) AddToHead(node *Node) {  
   node.Prev = l.Head  
   node.Next = l.Head.Next  
   l.Head.Next.Prev = node  
   l.Head.Next = node  
}  
func (l *LRUCache) RemoveTail() *Node {  
   node := l.Tail.Prev  
   l.RemoveNode(node)  
   return node  
}
```