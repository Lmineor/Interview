package tree

import (
	"fmt"
	"testing"
)

func TestReverse(t *testing.T){
	s := []int{1,2,3,4,5,6}
	reverse(s)
	fmt.Println(s)
}
