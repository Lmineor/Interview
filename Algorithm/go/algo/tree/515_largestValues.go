package tree

//给定一棵二叉树的根节点 root ，请找出该二叉树中每一层的最大值。

func largestValues(root *TreeNode) []int {
	var result []int
	if root == nil{
		return result
	}
	var level []*TreeNode

	var levelResult []int

	level = append(level, root)

	for len(level) > 0 {
		currentLen := len(level)
		levelResult = []int{}
		for i := 0; i < currentLen; i++ {
			node := level[0]
			level = level[1:]
			levelResult = append(levelResult, node.Val)
			if node.Left != nil {
				level = append(level, node.Left)
			}
			if node.Right != nil {
				level = append(level, node.Right)
			}
		}
		result = append(result, max(levelResult))
	}
	return result
}

func max(l []int) int {
	result := l[0]
	for _, num := range l {
		if result < num {
			result = num
		}
	}
	return result
}
