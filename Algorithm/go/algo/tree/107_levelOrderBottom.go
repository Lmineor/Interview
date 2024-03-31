package tree

func levelOrderBottom(root *TreeNode) [][]int {
	var result [][]int
	if root == nil {
		return result
	}
	var level []*TreeNode
	level = append(level, root) // 根节点入队
	var currentLevelResult []int
	var currentLevelLen int
	for len(level) > 0 {
		currentLevelResult = []int{}
		currentLevelLen = len(level)
		for i := 0; i < currentLevelLen; i++ {
			node := level[0]
			currentLevelResult = append(currentLevelResult, node.Val)
			level = level[1:]
			if node.Left != nil {
				level = append(level, node.Left)
			}
			if node.Right != nil {
				level = append(level, node.Right)
			}
		}
		result = append(result, currentLevelResult)
	}
	reverseS(result)
	return result
}


func reverseS(s [][]int){
	for i:= 0; i<len(s)/2; i++{
		s[i], s[len(s)-1-i] = s[len(s)-1-i],s[i]
	}
}