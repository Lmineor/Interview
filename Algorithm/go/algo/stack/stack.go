package stack

import "fmt"

func Entry(){
	//fmt.Println(computeExpression("1-2*3/3"))
	//fmt.Println(isValid("{[]}"))
	fmt.Println(evalRPN([]string{"2", "1", "+", "3", "*"}))
	fmt.Println(evalRPN([]string{"4", "13", "5", "/", "+"}))
	fmt.Println(evalRPN([]string{"10", "6", "9", "3", "+", "-11", "*", "/", "*", "17", "+", "5", "+"}))
}

