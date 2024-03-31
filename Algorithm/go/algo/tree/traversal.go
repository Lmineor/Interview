package tree

func preorderTraversal(root *Tree) []int {
	var f func(root *Tree)
	var result []int
	f = func(root *Tree) {
		if root != nil {
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
	f = func(root *Tree) {
		if root != nil {
			f(root.Left)
			f(root.Right)
			result = append(result, root.Val)
		}
	}
	f(root)
	return result
}

func inorderTraversal(root *Tree) []int {
	var f func(root *Tree)
	var result []int
	f = func(root *Tree) {
		if root != nil {
			f(root.Left)
			result = append(result, root.Val)
			f(root.Right)
		}
	}
	f(root)
	return result
}

func preorderTraversalByLevel(root *Tree) []int {
	var result []int
	level := Queue{ele: make([]*Tree, 0)}
	level.Push(root)

	for !level.Empty() {
		currentLevelNodeCount := len(level.ele)
		for i := 0; i < currentLevelNodeCount; i++ {
			node := level.Pop()
			result = append(result, node.Val)
			if node.Left != nil {
				level.Push(node.Left)
			}
			if node.Right != nil {
				level.Push(node.Right)
			}
		}
	}
	return result
}

func inorderTraversalByLevel(root *Tree) []int {
	var result []int
	level := Queue{ele: make([]*Tree, 0)}
	level.Push(root)

	for !level.Empty() {
		currentLevelNodeCount := len(level.ele)
		for i := 0; i < currentLevelNodeCount; i++ {
			node := level.Top()
			if node.
			result = append(result, node.Val)
			if node.Left != nil {
				level.Push(node.Left)
			}
			if node.Right != nil {
				level.Push(node.Right)
			}
		}
	}
	return result
}

type Queue struct {
	ele []*Tree
}

func (q *Queue) Push(node *Tree) {
	q.ele = append(q.ele, node)
}

func (q *Queue) Pop() *Tree {
	node := q.ele[0]
	q.ele = q.ele[1:]
	return node
}

func (q *Queue)Top() *Tree{
	return q.ele[0]
}

func (q *Queue) Empty() bool {
	return len(q.ele) == 0
}
