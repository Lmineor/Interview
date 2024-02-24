package link

//在链表类中实现这些功能：
//
//get(index)：获取链表中第 index 个节点的值。如果索引无效，则返回-1。
//addAtHead(val)：在链表的第一个元素之前添加一个值为 val 的节点。插入后，新节点将成为链表的第一个节点。
//addAtTail(val)：将值为 val 的节点追加到链表的最后一个元素。
//addAtIndex(index,val)：在链表中的第 index 个节点之前添加值为 val  的节点。如果 index 等于链表的长度，则该节点将附加到链表的末尾。
//如果 index 大于链表长度，则不会插入节点。如果index小于0，则在头部插入节点。
//deleteAtIndex(index)：如果索引 index 有效，则删除链表中的第 index 个节点。

type MyLinkedList struct {
	ele []int
}

func Constructor() MyLinkedList {
	return MyLinkedList{ele: make([]int, 0)}
}

func (this *MyLinkedList) length() int {
	return len(this.ele)
}

func (this *MyLinkedList) Get(index int) int {
	if this.length() <= index {
		return -1
	}
	return this.ele[index]
}

func (this *MyLinkedList) AddAtHead(val int) {
	this.ele = append([]int{val}, this.ele...)
}

func (this *MyLinkedList) AddAtTail(val int) {
	this.ele = append(this.ele, val)

}

func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index <= -1 {
		return
	}
	if index == this.length() {
		this.ele = append(this.ele, val)
		return
	}
	if index > this.length() {
		return
	}
	after := make([]int, len(this.ele[index:]))
	copy(after, this.ele[index:])
	this.ele = append(this.ele[:index], val)
	this.ele = append(this.ele, after...)
}

func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index <= -1 {
		return
	}
	if this.length() <= index {
		return
	}
	this.ele = append(this.ele[:index], this.ele[index+1:]...)
}
