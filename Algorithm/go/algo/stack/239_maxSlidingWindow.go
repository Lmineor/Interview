package stack

func maxSlidingWindow(nums []int, k int) []int {
	left := 0
	right := k
	result := make([]int, 0)
	currentMax, index := max(nums[0:k])
	result = append(result, currentMax)
	for right < len(nums)-1 {
		right++
		left++
		if nums[right] >= currentMax {
			index = right
		}
		if index >= left {
			result = append(result, nums[index])
		} else {
			currentMax, index = max(nums[left:right])
			result = append(result, currentMax)
		}
	}
	return result
}

type MyQ struct {
	ele []int
}

func (q *MyQ) Push(value int) {
	//push(value)：如果push的元素value大于入口元素的数值，那么就将队列入口的元素弹出，直到push元素的数值小于等于队列入口元素的数值为止
	for !q.Empty() && q.Back() < value {
		q.ele = q.ele[:len(q.ele)-1]
	}
	q.ele = append(q.ele, value)
}

func (q *MyQ) Back() int {
	return q.ele[len(q.ele)-1]
}
func (q *MyQ) Front() int {
	return q.ele[0]
}
func (q *MyQ) Pop(value int) int {
	//pop(value)：如果窗口移除的元素value等于单调队列的出口元素，那么队列弹出元素，否则不用任何操作
	if !q.Empty() && q.Front() == value {
		q.ele = q.ele[1:]
	}
}

func (q *MyQ) Empty() bool {
	return len(q.ele) == 0
}
