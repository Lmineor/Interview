package stack

type MyStack struct {
	InQueue  []int
	OutQueue []int
}

func Constructor() MyStack {
	return MyStack{
		InQueue:  make([]int, 0),
		OutQueue: make([]int, 0),
	}
}

func (this *MyStack) Push(x int) {
	this.InQueue = append(this.InQueue, x)
}

func (this *MyStack) Pop() int {
	inLen, outLen := len(this.InQueue), this.OutQueue
	if outLen == 0{
		if inLen == 0{
			return -1
		}
		for i:= inLen-1; i>=0; i--{
			this.OutQueue = append(this.OutQueue, this.InQueue[i])
		}
		this.InQueue = []int{}
		outLen = len(this.OutQueue)
	}
	val := this.OutQueue[outLen-1]
}

func (this *MyStack) Top() int {

<-1,2,3,4,5,6,7<-
<-1,2,3,4,5,6,7,
}

func (this *MyStack) Empty() bool {

}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
