package tree

//给定一棵二叉树，想象自己站在它的右侧，按照从顶部到底部的顺序，返回从右侧所能看到的节点值。

func rightSideView(root *TreeNode) []int {
	var result []int
	if root == nil {
		return result
	}
	var level []*TreeNode
	var currentLevel []int
	var levelResult [][]int

	level = append(level, root)

	for len(level) > 0 {
		currentLevelLen := len(level)
		currentLevel = []int{}
		for i := 0; i < currentLevelLen; i++ {
			node := level[0]
			level = level[1:]
			currentLevel = append(currentLevel, node.Val)
			if node.Left != nil {
				level = append(level, node.Left)
			}
			if node.Right != nil {
				level = append(level, node.Right)
			}
		}
		levelResult = append(levelResult, currentLevel)
	}
	for _, l := range levelResult {
		result = append(result, l[len(l)-1])
	}
	return result
}
