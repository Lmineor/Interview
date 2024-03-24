package stack

import (
	"sort"
)

//给定一个非空的整数数组，返回其中出现频率前 k 高的元素。
//
//示例 1:
//
//输入: nums = [1,1,1,2,2,3], k = 2
//输出: [1,2]
//示例 2:
//
//输入: nums = [1], k = 1
//输出: [1]

func topKFrequent(nums []int, k int) []int {
	numMap := make(map[int]int)
	for _, num:= range nums{
		numMap[num]++
	}
	result := make([]int, 0)
	for key, _ := range numMap{
		result = append(result, key)
	}
	sort.Slice(result, func(i, j int) bool {
		return numMap[result[i]] > numMap[result[j]]
	})
	return result[:k]
}

