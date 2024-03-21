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

func max(nums []int)(index int, value int){
	result := nums[0]
	index = 0
	for k,v := range nums{
		if v > result{
			result = v
			index = k
		}
	}
	return index, result
}