package array

import "fmt"

func Entry(){
	fmt.Println(spiralOrder(generateMatrix(4)))
	//a := [][]int{{1,2,3,4}}
	a := [][]int{{1,2,3,4},{5,6,7,8}, {9,10,11,12}}
	fmt.Println(spiralOrder(a))
}