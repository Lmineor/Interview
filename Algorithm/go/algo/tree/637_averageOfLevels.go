package tree

//给定一个非空二叉树的根节点 root , 以数组的形式返回每一层节点的平均值。与实际答案相差 10-5 以内的答案可以被接受。

func averageOfLevels(root *TreeNode) []float64 {
	var result []float64
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
			currentLevel = append(currentLevel, node.Val)
			level = level[1:]
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
		result = append(result, ave(l))
	}
	return result
}

func ave(l []int) float64 {
	sum := 0
	for _, i := range l {
		sum = sum + i
	}
	return float64(sum) / float64(len(l))
}
