package stack

type MyStack struct {
	q1 []int
	q2 []int
}

func Constructor1() MyStack {
	return MyStack{
		q1: make([]int, 0),
		q2: make([]int, 0),
	}
}

func (this *MyStack) Push(x int) {
	q1Len := len(this.q1)
	if q1Len != 0 {
		this.q1 = append(this.q1, x)
	} else {
		this.q2 = append(this.q2, x)
	}
}

func (this *MyStack) Pop() int {
	var lastOne = -1
	for len(this.q1) != 0 {
		lastOne = this.q1[0]
		if len(this.q1) == 1 {
			this.q1 = []int{}
			return lastOne
		}
		this.q2 = append(this.q2, lastOne)
		this.q1 = this.q1[1:]
	}
	for len(this.q2) != 0 {
		lastOne = this.q2[0]
		if len(this.q2) == 1 {
			this.q2 = []int{}
			return lastOne
		}
		this.q1 = append(this.q1, lastOne)
		this.q2 = this.q2[1:]
	}
	return lastOne
}

func (this *MyStack) Top() int {
	var lastOne = -1
	if len(this.q1) !=0{
		for len(this.q1) != 0 {
			lastOne = this.q1[0]
			this.q2 = append(this.q2, lastOne)
			this.q1 = this.q1[1:]
		}
		return lastOne
	}else{
		for len(this.q2) != 0 {
			lastOne = this.q2[0]
			this.q1 = append(this.q1, lastOne)
			this.q2 = this.q2[1:]
		}
		return lastOne
	}
}

func (this *MyStack) Empty() bool {
	return len(this.q1) == 0 && len(this.q2) == 0
}

/**
 * Your MyStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * param_2 := obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.Empty();
 */
