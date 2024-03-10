package str

import "fmt"

func twoSum(target int, data []int) (int, int) {
	tmp := make(map[int]int, 0)
	for i := 0; i < len(data); i++ {
		res := target - data[i]
		if in(res, tmp) {
			return i, tmp[res]
		}
		tmp[data[i]] = i
	}
	return 0, 0
}

func in(num int, amap map[int]int) bool {
	_, ok := amap[num]
	return ok
}

func Entry() {
	//fmt.Println(twoSum(10, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}))
	fmt.Println(strStr("mississippi", "issip"))
	fmt.Println(strStr("a", "a"))
}
