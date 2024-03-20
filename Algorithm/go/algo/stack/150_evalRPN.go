package stack

import "strconv"

const (
	plus  = "+"
	minus = "-"
	cheng = "*"
	chu   = "/"
)

func evalRPN(tokens []string) int {
	i := 0
	for len(tokens) > 1 {
		if !isNum(tokens[i]){
			leftNum := tokens[i-2]
			rightNum := tokens[i-1]
			op := tokens[i]
			val := cal(leftNum,rightNum,op)
		}
	}
}

func cal(leftNum, rightNum, op string)int{
	l,_ := strconv.Atoi(leftNum)
	r, _:= strconv.Atoi(rightNum)
	switch op {
	case plus:
		return l+r
	case minus:
		return l-r
	case cheng:
		return l*r
	case chu:
		return l/r
	}
}

func isNum(c string) bool {
	nums := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	for _, num := range nums {
		if c == num {
			return true
		}
	}
	return false
}
