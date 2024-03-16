package stack

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
//
//有效字符串需满足：
//
//左括号必须用相同类型的右括号闭合。
//左括号必须以正确的顺序闭合。
//注意空字符串可被认为是有效字符串。

func isValid(s string) bool {
	var a = "("
	var b = ")"
	var c = "["
	var d = "]"
	var e = "{"
	var f = "}"

	stack := make([]string, 0)
	for _, ss := range s {
		switch string(ss) {
		case a, c, e:
			stack = append(stack, string(ss))
		case b:
			if len(stack) == 0 || stack[len(stack)-1] != a {
				return false
			}
			stack = stack[:len(stack)-1]
		case d:
			if len(stack) == 0 || stack[len(stack)-1] != c {
				return false
			}
			stack = stack[:len(stack)-1]
		case f:
			if len(stack) == 0 || stack[len(stack)-1] != e {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}
