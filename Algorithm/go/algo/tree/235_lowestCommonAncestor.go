package tree

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val   int
 *     Left  *TreeNode
 *     Right *TreeNode
 * }
 */

var stack []*TreeNode
var result []*TreeNode

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	stack = make([]*TreeNode, 0)
	var pStack []*TreeNode
	dfsSearch(root, p)
	pStack = result
	stack = stack[:0]

	var qStack []*TreeNode
	dfsSearch(root, q)
	qStack = result

	var targetNode *TreeNode
	var smaller int
	if len(pStack) < len(qStack) {
		smaller = len(pStack)
	} else {
		smaller = len(qStack)
	}
	for i := 0; i < smaller; i++ {
		if qStack[i] == pStack[i] {
			targetNode = qStack[i]
		}
	}
	return targetNode
}

func dfsSearch(root *TreeNode, target *TreeNode) {
	if root == nil {
		return
	}
	if root == target {
		stack = append(stack, root)
		result = make([]*TreeNode, len(stack))
		copy(result, stack)
		return
	}
	stack = append(stack, root)
	dfsSearch(root.Left, target)
	dfsSearch(root.Right, target)
	stack = stack[:len(stack)-1]
}
