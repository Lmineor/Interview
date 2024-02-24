package array

//给定一个 n 个元素有序的（升序）整型数组 nums 和一个目标值 target  ，写一个函数搜索 nums 中的 target，如果目标值存在返回下标，否则返回 -1。

func BinarySearch(list []int, target int) int {
	end := len(list) - 1
	start := 0
	var mid int
	for start <= end {
		mid = (start + end) / 2
		current := list[mid]
		if current == target {
			return mid
		} else if current < target {
			start = mid + 1
		} else if current > target {
			end = mid - 1
		}
	}
	return -1
}
