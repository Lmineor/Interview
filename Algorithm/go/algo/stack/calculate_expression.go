package stack
//
//import (
//	"strconv"
//)
//
//const (
//	plus     = '+'
//	minus    = '-'
//	multiply = '*'
//	divide   = '/'
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
//	if s.Empty() {
//		panic("pop from empty stack")
//	}
//	element := s.ele[0]
//	s.ele = s.ele[1:]
//	return element
//}
//
//func (s *Stack) Pop() interface{} {
//	if s.Empty() {
//		panic("pop from empty stack")
//	}
//	element := s.ele[len(s.ele)-1]
//	s.ele = s.ele[:len(s.ele)-1]
//	return element
//}
//
//func (s *Stack) GetTop() interface{} {
//	if s.Empty() {
//		return rune(0)
//	}
//	return s.ele[len(s.ele)-1]
//}
//
//func (s *Stack) Push(ele interface{}) {
//	s.ele = append(s.ele, ele)
//}
//
//func (s *Stack) PushLeft(ele interface{}) {
//	s.ele = append([]interface{}{ele}, s.ele...)
//}
//
//func computeExpression(exp string) int {
//	numStack := Stack{
//		ele: make([]interface{}, 0),
//	}
//	operateStack := Stack{
//		ele: make([]interface{}, 0),
//	}
//	for _, c := range exp {
//		switch c {
//		case plus, minus, multiply, divide:
//			operateStack.Push(c)
//		default:
//			num, _ := strconv.Atoi(string(c))
//			op := operateStack.GetTop().(rune)
//
//			if op == multiply || op == divide {
//				op = operateStack.Pop().(rune)
//				numLeft := numStack.Pop().(int)
//				numStack.Push(compute(numLeft, num, op))
//			} else {
//				numStack.Push(num)
//			}
//		}
//	}
//	for !operateStack.Empty() {
//		numLeft := numStack.PopLeft().(int)
//		numRight := numStack.PopLeft().(int)
//		op := operateStack.PopLeft().(rune)
//		numStack.PushLeft(compute(numLeft, numRight, op))
//	}
//	return numStack.PopLeft().(int)
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
