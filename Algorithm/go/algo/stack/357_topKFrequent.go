package stack

import "sort"

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
	cache := make(map[int]int, 0)
	for _, num := range nums{
		cache[num]++
	}
	count := make([]int, 0)
	for _, c := range cache{
		count = append(count, c)
	}
	sort.Ints(count)
	count = count[len(count)-2:]
	result := make([]int, 0)
}

