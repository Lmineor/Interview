package stack

//使用栈实现队列的下列操作：
//
//push(x) -- 将一个元素放入队列的尾部。
//pop() -- 从队列首部移除元素。
//peek() -- 返回队列首部的元素。
//empty() -- 返回队列是否为空。

type MyQueue struct {
	StackIn  []int
	StackOut []int
}

func Constructor() MyQueue {
	return MyQueue{
		StackOut: make([]int, 0),
		StackIn: make([]int,0),
	}
}

func (this *MyQueue) Push(x int) {
	this.StackIn = append(this.StackIn, x)

}

func (this *MyQueue) Pop() int {
	inLen,outLen := len(this.StackIn), len(this.StackOut)
	if outLen== 0{
		if inLen == 0{
			return -1
		}
		for i:=inLen-1;i>=0;i--{
			this.StackOut = append(this.StackOut, this.StackIn[i])
		}
		outLen = len(this.StackOut)
		this.StackIn = []int{}
	}
	val := this.StackOut[outLen-1]
	this.StackOut = this.StackOut[:outLen-1]
	return val
}

func (this *MyQueue) Peek() int {
	inLen,outLen := len(this.StackIn), len(this.StackOut)
	if outLen== 0{
		if inLen == 0{
			return -1
		}
		for i:=inLen-1;i>=0;i--{
			this.StackOut = append(this.StackOut, this.StackIn[i])
		}
		outLen = len(this.StackOut)
		this.StackIn = []int{}
	}
	return this.StackOut[len(this.StackOut)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.StackIn) == 0 && len(this.StackOut)==0

}

/**
 * Your MyQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Peek();
 * param_4 := obj.Empty();
 */
