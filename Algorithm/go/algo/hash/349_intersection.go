package hash


//题意：给定两个数组，编写一个函数来计算它们的交集。
//
//349. 两个数组的交集
//
//说明： 输出结果中的每个元素一定是唯一的。 我们可以不考虑输出结果的顺序。
//
//#

func intersection2(nums1 []int, nums2 []int) []int {
	var result []int
	resultMap := make(map[int]struct{}, 0)
	bucket := make(map[int]struct{},0)
	for _, num := range nums1{
		bucket[num] = struct{}{}
	}
	for _, num:= range nums2{
		if _, ok := bucket[num]; ok{
			resultMap[num] = struct{}{}
		}
	}
	for k, _:= range resultMap{
		result = append(result, k)
	}
	return result
}

func intersection(nums1 []int, nums2 []int) []int {
	var result []int
	resultMap := make(map[int]int, 0)
	bucket := make(map[int]int,0)
	for _, num := range nums1{
		if _, ok:= bucket[num];!ok{
			bucket[num] = 1
		}else{
			bucket[num] = bucket[num]+1
		}

	}
	for _, num:= range nums2{
		if _, ok := bucket[num]; ok{
			resultMap[num] = resultMap[num]+1
		}else{
			resultMap[num] = 1
		}
	}
	for k, _:= range resultMap{
		result = append(result, k)
	}
	return result
}
