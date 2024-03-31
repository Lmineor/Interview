package tree

import "fmt"

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

func TraversalByLevel(root *Tree) []int {
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

func preorderTraversalByLevel(root *Tree) []int {
	var result []int
	level := Queue{ele: make([]*Tree, 0)}
	level.Push(root)

	for !level.Empty() {
		node := level.Pop()
		result = append(result, node.Val)
		if node.Right != nil {
			level.Push(node.Right)
		}
		if node.Left != nil {
			level.Push(node.Left)
		}
	}
	return result
}

// 非迭代中序遍历二叉树
func inorderTraversalNoRec(root *Tree) []int {
	var result []int
	if root == nil {
		return result
	}
	p := root
	level := Queue{ele: make([]*Tree, 0)}
	for !level.Empty() || p != nil {
		if p != nil {
			level.Push(p)
			p = p.Left
		} else {
			p = level.PopRight()
			fmt.Println(p.Val)
			result = append(result, p.Val)
			p = p.Right
		}
	}

	return result
}

func postOrderTraversalNoRec(root *Tree) []int {
	var result []int
	level := Queue{ele: make([]*Tree, 0)}
	level.Push(root)

	for !level.Empty() {
		node := level.Pop()
		result = append(result, node.Val)
		if node.Left != nil {
			level.Push(node.Left)
		}
		if node.Right != nil {
			level.Push(node.Right)
		}
	}
	reverse(result)
	return result
}

func reverse(s []int){
	for i:= 0; i<len(s)/2; i++{
		s[i], s[len(s)-1-i] = s[len(s)-1-i],s[i]
	}
}

type Queue struct {
	ele []*Tree
}

func (q *Queue) Push(node *Tree) {
	q.ele = append(q.ele, node)
}

func (q *Queue) Pop() *Tree {
	node := q.ele[len(q.ele)-1]
	q.ele = q.ele[:len(q.ele)-1]
	return node
}

func (q *Queue) Top() *Tree {
	return q.ele[len(q.ele)-1]
}

func (q *Queue) PopRight() *Tree {
	node := q.ele[len(q.ele)-1]
	q.ele = q.ele[:len(q.ele)-1]
	return node
}

func (q *Queue) Empty() bool {
	return len(q.ele) == 0
}
