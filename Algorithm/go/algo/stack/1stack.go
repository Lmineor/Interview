package stack

//
//import "fmt"
//
//const (
//	leftParenthesis  = '('
//	rightParenthesis = ')'
//	plus             = '+'
//	minus            = '-'
//	multiply         = '*'
//	divide           = '/'
//)
//
//type Stack struct {
//	ele []interface{}
//}
//
//func (s *Stack) Empty() bool {
//	return len(s.ele) == 0
//}
//func (s *Stack) PopLeft() interface{} {
//	if !s.Empty() {
//		element := s.ele[len(s.ele)-1]
//		s.ele = s.ele[:len(s.ele)-1]
//		return element
//	}
//	return 0
//}
//
//func (s *Stack) Push(ele interface{}) {
//	s.ele = append(s.ele, ele)
//}
//
//func computeExpression(exp string) int {
//	var result int
//	numStack := Stack{
//		ele: make([]interface{}, 0),
//	}
//	cStack := Stack{
//		ele: make([]interface{}, 0),
//	}
//	for _, c := range exp {
//		switch c {
//		case leftParenthesis, rightParenthesis, plus, minus, multiply, divide:
//			cStack.Push(c)
//		default:
//			numStack.Push(c)
//		}
//	}
//	for !numStack.Empty() {
//
//	}
//	for _, c := range exp {
//		switch c {
//		case leftParenthesis:
//			cStack.Push(leftParenthesis)
//		case rightParenthesis:
//			op := cStack.PopLeft().(rune)
//			cStack.PopLeft() // (
//			numLeft := numStack.PopLeft().(int)
//			numRight := numStack.PopLeft().(int)
//			numStack.Push(compute(numLeft, numRight, op))
//		case plus:
//			cStack.Push(plus)
//		case minus:
//			cStack.Push(minus)
//		case multiply:
//			numLeft := numStack.PopLeft().(int)
//			numRight := numStack.PopLeft().(int)
//			numStack.Push(numLeft * numRight)
//		case divide:
//			numLeft := numStack.PopLeft().(int)
//			numRight := numStack.PopLeft().(int)
//			numStack.Push(numLeft / numRight)
//		default:
//			numStack.Push(int(c))
//		}
//	}
//	for !numStack.Empty() {
//		numLeft := numStack.PopLeft().(int)
//		numRight := numStack.PopLeft().(int)
//		op := cStack.PopLeft().(rune)
//		result += compute(numLeft, numRight, op)
//	}
//	return result
//}
//
//func compute(numLeft, numRight int, op rune) int {
//	switch op {
//	case plus:
//		return numLeft + numRight
//	case minus:
//		return numLeft - numRight
//	case multiply:
//		return numLeft * numRight
//	case divide:
//		return numLeft / numRight
//	}
//	return 0
//}
//
//func isOP(c rune) bool {
//	switch c {
//	case leftParenthesis, rightParenthesis, plus, minus, multiply, divide:
//		return true
//	}
//	return false
//}
//func Entry() {
//	fmt.Println(computeExpression("1+2+3+4"))
//}
