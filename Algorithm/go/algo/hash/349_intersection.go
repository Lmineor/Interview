package hash

//题意：给定两个数组，编写一个函数来计算它们的交集。
//
//349. 两个数组的交集
//
//说明： 输出结果中的每个元素一定是唯一的。 我们可以不考虑输出结果的顺序。
//
//# 350.
//给你两个整数数组 nums1 和 nums2 ，请你以数组形式返回两数组的交集。返回结果中每个元素出现的次数，应与元素在两个数组中都出现的次数一致（如果出现次数不一致，则考虑取较小值）。可以不考虑输出结果的顺序。
//
//
//
//示例 1：
//
//输入：nums1 = [1,2,2,1], nums2 = [2,2]
//输出：[2,2]
//示例 2:
//
//输入：nums1 = [4,9,5], nums2 = [9,4,9,8,4]
//输出：[4,9]

//349
func intersection2(nums1 []int, nums2 []int) []int {
	var result []int
	resultMap := make(map[int]struct{}, 0)
	bucket := make(map[int]struct{}, 0)
	for _, num := range nums1 {
		bucket[num] = struct{}{}
	}
	for _, num := range nums2 {
		if _, ok := bucket[num]; ok {
			resultMap[num] = struct{}{}
		}
	}
	for k, _ := range resultMap {
		result = append(result, k)
	}
	return result
}


//350
func intersection(nums1 []int, nums2 []int) []int {
	var result []int
	resultMap := make(map[int]int, 0)
	bucket := make(map[int]int, 0)
	// 统计nums1中的数字的频数
	for _, num := range nums1 {
		if _, ok := bucket[num]; !ok {
			bucket[num] = 1
		} else {
			bucket[num] = bucket[num] + 1
		}
	}

	// 重复的数字在nums2中出现的次数
	for _, num := range nums2 {
		if _, ok := bucket[num]; ok {
			if _, ok1 := resultMap[num]; ok1 {
				resultMap[num] = resultMap[num] + 1
			} else {
				resultMap[num] = 1
			}
		}

	}

	for k, v := range resultMap {
		minNum := 1
		if v <= bucket[k] {
			minNum = v
		} else {
			minNum = bucket[k]
		}
		for i := 0; i < minNum; i++ {
			result = append(result, k)
		}
	}
	return result
}
