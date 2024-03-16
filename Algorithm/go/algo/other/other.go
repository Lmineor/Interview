package other

import "fmt"

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n > 1 {
		return fib(n-1) + fib(n-2)
	}
	return 0
}

func Entry() {

	//fmt.Println(mySqrt(8))
	lru :=Constructor()
	fmt.Println(lru.Get(123))
	for i:=0; i<15;  i++{
		lru.Put(i, i)
	}
	fmt.Println(lru.Get(14))
	fmt.Println(lru.Cache)
}
