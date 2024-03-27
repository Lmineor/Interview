package tree

import "fmt"

func PrintInOrder(t *Tree) {
	if t != nil {
		PrintInOrder(t.Left)
		fmt.Println(t.Val)
		PrintInOrder(t.Right)
	}
}

func PrintPreOrder(t *Tree) {
	if t != nil {
		fmt.Println(t.Val)
		PrintPreOrder(t.Left)
		PrintPreOrder(t.Right)
	}
}

func PrintPostOrder(t *Tree) {
	if t != nil {
		PrintPostOrder(t.Left)
		PrintPostOrder(t.Right)
		fmt.Println(t.Val)
	}
}
func PrintInLevel(t *Tree) [][]int {
	var level []*Tree
	var levelResult [][]int
	var currentLevelResult []int
	if t == nil {
		return [][]int{}
	}
	cur := t
	level = append(level, cur) // 根节点入
	for len(level) > 0 {
		currentLevelLength := len(level)
		for i := 0; i < currentLevelLength; i++ {
			popped := level[i]
			currentLevelResult = append(currentLevelResult, popped.Val)
			if popped.Left != nil {
				level = append(level, popped.Left)
			}
			if popped.Right != nil {
				level = append(level, popped.Right)
			}
		}
		level = level[currentLevelLength:]
		levelResult = append(levelResult, currentLevelResult)
		currentLevelResult = []int{}
	}
	return levelResult
}
func preorderTraversal(root *Tree) []int {
	var f func(root *Tree)
	var result []int
	f = func (root *Tree){
		if root != nil{
			result = append(result, root.Val)
			f(root.Left)
			f(root.Right)
		}
	}
	f(root)
	return result
}

func postorderTraversal(root *Tree) []int {
	var f func(root *Tree)
	var result []int
	f = func (root *Tree){
		if root != nil{
			f(root.Left)
			f(root.Right)
			result = append(result, root.Val)
		}
	}
	f(root)
	return result
}


func inorderTraversal(root *Tree)[]int{
	var f func(root *Tree)
	var result []int
	f = func(root *Tree){
		if root != nil{
			f(root.Left)
			result = append(result, root.Val)
			f(root.Right)
		}
	}
	f(root)
	return result
}