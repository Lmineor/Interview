package stack


func maxSlidingWindow(nums []int, k int) []int {
	left:=0
	right := k
	result := make([]int, 0)
	currentMax := max(nums[0:k])
	result = append(result, currentMax)
	for right < len(nums){
		right++
		if nums[right]>currentMax{
			result = append(result, nums[right])
		}
	}
}

func max(nums []int)int{
	result := nums[0]
	for _,v := range nums{
		if v > result{
			result = v
		}
	}
	return result
}