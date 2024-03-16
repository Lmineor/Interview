package stack

//使用栈实现队列的下列操作：
//
//push(x) -- 将一个元素放入队列的尾部。
//pop() -- 从队列首部移除元素。
//peek() -- 返回队列首部的元素。
//empty() -- 返回队列是否为空。

type MyQueue struct {
	StackIn []int
	StackOut []int
}

func Constructor() MyQueue {

}

func (this *MyQueue) Push(x int) {
	this.StackIn = append(this.StackIn, x)
	this.StackOut = 

}

func (this *MyQueue) Pop() int {

}

func (this *MyQueue) Peek() int {

}

func (this *MyQueue) Empty() bool {

}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
