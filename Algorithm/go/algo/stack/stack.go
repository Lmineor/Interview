package stack

import "fmt"

func Entry(){
	//fmt.Println(computeExpression("1-2*3/3"))
	s := Constructor1()
	//fmt.Println(s.Pop())
	//fmt.Println(s.Peek())
	s.Push(1)
	s.Push(2)
	//fmt.Println(s.Peek())
	s.Push(3)
	s.Push(4)
	fmt.Println(s.Pop())
	s.Push(5)
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())
	fmt.Println(s.Pop())

}
