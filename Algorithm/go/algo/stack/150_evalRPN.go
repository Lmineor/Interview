package stack

import (
	"strconv"
)

const (
	plus  = "+"
	minus = "-"
	cheng = "*"
	chu   = "/"
)

func evalRPN(tokens []string) int {
	numStack := make([]int, 0)
	for _, token := range tokens{
		switch token {
		case plus, minus,cheng,chu:
			rightNum := numStack[len(numStack)-1]
			leftNum := numStack[len(numStack)-2]
			numStack = numStack[:len(numStack)-2]
			val := cal(leftNum, rightNum, token)
			numStack = append(numStack, val)
		default:
			num, _:= strconv.Atoi(token)
			numStack  = append(numStack, num)
		}
	}
	return numStack[0]
}

func cal(l, r int, op string) int {
	switch op {
	case plus:
		return l + r
	case minus:
		return l - r
	case cheng:
		return l * r
	default:
		return l / r
	}
}
