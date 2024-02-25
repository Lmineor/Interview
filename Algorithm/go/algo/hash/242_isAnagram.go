package hash

//给定两个字符串 s 和 t ，编写一个函数来判断 t 是否是 s 的字母异位词。
//
//示例 1: 输入: s = "anagram", t = "nagaram" 输出: true
//
//示例 2: 输入: s = "rat", t = "car" 输出: false
//
//说明: 你可以假设字符串只包含小写字母。
//
//

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sMap := make(map[rune]int, 0)
	for _, c := range s {
		if _, ok := sMap[c]; ok {
			sMap[c] = sMap[c] + 1
		} else {
			sMap[c] = 1
		}
	}
	for _, c := range t {
		if _, ok := sMap[c]; ok {
			sMap[c] = sMap[c] - 1
		} else {
			return false
		}
	}
	for _,v := range sMap{
		if v != 0{
			return false
		}
	}
	return true
}
