package array

//给你一个按 非递减顺序 排序的整数数组 nums，返回 每个数字的平方 组成的新数组，要求也按 非递减顺序 排序。
//
//示例 1：
//
//输入：nums = [-4,-1,0,3,10]
//输出：[0,1,9,16,100]
//解释：平方后，数组变为 [16,1,0,9,100]，排序后，数组变为 [0,1,9,16,100]
//示例 2：
//
//输入：nums = [-7,-3,2,3,11]
//输出：[4,9,9,49,121]

func sortedSquares(nums []int) []int {
	result := make([]int, len(nums), len(nums))
	startIndex := 0
	endIndex := len(nums) - 1
	for i := len(nums) - 1; i >= 0; i-- {
		startSqu := nums[startIndex] * nums[startIndex]
		endSqu := nums[endIndex] * nums[endIndex]
		if startSqu < endSqu {
			result[i] = endSqu
			endIndex--
		} else {
			result[i] = startSqu
			startIndex++
		}
	}
	return result
}
