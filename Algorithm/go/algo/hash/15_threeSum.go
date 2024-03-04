package hash


//给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有满足条件且不重复的三元组。
//
//注意： 答案中不可以包含重复的三元组。
//
//示例：
//
//给定数组 nums = [-1, 0, 1, 2, -1, -4]，
//
//满足要求的三元组集合为： [ [-1, 0, 1], [-1, -1, 2] ]
//
//#


func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	for index, num := range nums{
		otherSlice := nums[:index]
		otherSlice = append(otherSlice, nums[index:]...)
		if num1, num2, find := twoSum1(otherSlice, -num); find{
			anAnswer := []int{num, num1, num2}
			result = append(result, anAnswer)
		}
	}
	return result
}

func twoSum1(nums []int, target int)(nums1, nums2 int, find bool){
	record := make(map[int]int)
	for _, num := range nums{
		if _, ok := record[target-num];ok {
			return num, target-num, true
		}else{
			record[num]=1
		}
	}
	return 0, 0, false
}