package stack


func maxSlidingWindow(nums []int, k int) []int {
	left:=0
	right := k
	result := make([]int, 0)
	currentMax, index := max(nums[0:k])
	result = append(result, currentMax)
	for right < len(nums)-1{
		right++
		left++
		if nums[right]>=currentMax{
			index = right
		}
		if index >=left{
			result = append(result, nums[index])
		}else{
			currentMax, index = max(nums[left:right])
			result = append(result, currentMax)
		}
	}
	return result
}

type MyQueue struct {
	ele []int
}

func (q *MyQueue)Pop()int{
	//pop(value)：如果窗口移除的元素value等于单调队列的出口元素，那么队列弹出元素，否则不用任何操作
	//push(value)：如果push的元素value大于入口元素的数值，那么就将队列入口的元素弹出，直到push元素的数值小于等于队列入口元素的数值为止
	
}