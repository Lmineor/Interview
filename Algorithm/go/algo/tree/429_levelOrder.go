package tree

// 给定一个 N 叉树，返回其节点值的层序遍历。（即从左到右，逐层遍历）。
//
//树的序列化输入是用层序遍历，每组子节点都由 null 值分隔（参见示例）。

type Node struct {
	Val      int
	Children []*Node
}

func levelOrderN(root *Node) [][]int {
	var result [][]int
	var level []*Node

	if root == nil{
		return result
	}
	var levelResult []int
	level  = append(level, root)
	for len(level)> 0{
		levelResult = []int{}
		currentLen := len(level)
		for i:= 0; i< currentLen; i++{
			node := level[0]
			levelResult = append(levelResult, node.Val)
			level = level[1:]
			for _, item := range node.Children{
				if item != nil{
					level = append(level, item)
				}

			}
		}
		result = append(result, levelResult)
	}
	return result
}
