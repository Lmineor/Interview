package array

// 给定一个含有 n 个正整数的数组和一个正整数 s ，找出该数组中满足其和 ≥ s 的长度最小的 连续 子数组，并返回其长度。如果不存在符合条件的子数组，返回 0。
//
// 示例：
//
// 输入：s = 7, nums = [2,3,1,2,4,3] 输出：2 解释：子数组 [4,3] 是该条件下的长度最小的子数组。

func minSubArrayLen(target int, nums []int) int {
	l := len(nums)
	sum := 0
	j := 0
	result := l + 1
	for i := 0; i < l; i++ {
		sum += nums[i]
		for sum >= target {
			subLength := i - j + 1
			if subLength < result {
				result = subLength
			}
			sum -= nums[j]
			j++
		}
	}
	if result == l+1 {
		return 0
	} else {
		return result
	}

}
