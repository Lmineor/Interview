package tree

import "fmt"

func PrintInOrder(t *TreeNode) {
	if t != nil {
		PrintInOrder(t.Left)
		fmt.Println(t.Val)
		PrintInOrder(t.Right)
	}
}

func PrintPreOrder(t *TreeNode) {
	if t != nil {
		fmt.Println(t.Val)
		PrintPreOrder(t.Left)
		PrintPreOrder(t.Right)
	}
}

func PrintPostOrder(t *TreeNode) {
	if t != nil {
		PrintPostOrder(t.Left)
		PrintPostOrder(t.Right)
		fmt.Println(t.Val)
	}
}
func PrintInLevel(t *TreeNode) [][]int {
	var level []*TreeNode
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
