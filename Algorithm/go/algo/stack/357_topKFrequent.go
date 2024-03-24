package stack

import (
	"fmt"
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
	var count []int
	for _, c := range numMap {
		count = append(count, c)
	}
	fmt.Println(count)
	sort.Ints(count)
	count = count[len(count)-k:]

	var result []int
	for _, c := range count{
		result = append(result, countMap[c])
	}
	return result
}

