package main

//func main() {
//	//array.Entry()
//	hash.Entry()
//}


import "fmt"

func main() {
	fmt.Println("return:", test()) // defer 和 return之间的顺序是先返回值, i=0，后defer
}

func test()  (i int) { //这里返回值没有命名
	defer func() {
		i++
		fmt.Println("defer1", i) //作为闭包引用的话，则会在defer函数执行时根据整个上下文确定当前的值。i=2
	}()
	defer func() {
		i++
		fmt.Println("defer2", i) //作为闭包引用的话，则会在defer函数执行时根据整个上下文确定当前的值。i=1
	}()
	return i
}
