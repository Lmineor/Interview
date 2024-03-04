package hash

import "sort"

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

//思路 先排序，后双指针


func threeSum(nums []int) [][]int {
	result := make([][]int, 0)
	if len(nums)<3{
		return result
	}
	sort.Ints(nums)
	for i:=0; i<len(nums)-2;i++{
		n1 := nums[i]
		if n1 > 0{
			break
		}
		if i>0&& n1== nums[i-1]{
			continue
		}
		l, r:= i+1, len(nums)-1
		for l< r{
			n2, n3 := nums[l], nums[r]
			if n1+n2+n3==0{
				result = append(result, []int{n1, n2, n3})
				for l< r&&nums[l] == n2{
					l++
				}
				for l<r && nums[r] == n3{
					r--
				}
			}else if n1+n2+n3<0{
				l++
			}else{
				r--
			}
		}
	}
	return result
}
