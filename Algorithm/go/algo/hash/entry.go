package hash

import "fmt"

func Entry() {
	//fmt.Println(isAnagram("cat", "car"))
	//fmt.Println(intersection([]int{4,9,5}, []int{9,4,9,8,4}))
	//fmt.Println(isHappy(19))
	nums1 := []int{-1,-1}
	nums2 := []int{-1,1}
	nums3 := []int{-1,1}
	nums4 := []int{1,-1}
	fmt.Println(fourSumCount(nums1, nums2, nums3, nums4))
}
