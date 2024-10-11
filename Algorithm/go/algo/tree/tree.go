package tree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Build() *TreeNode {
	node6 := &TreeNode{
		Val: 6,
	}
	node2 := &TreeNode{
		Val: 2,
	}
	node8 := &TreeNode{
		Val: 8,
	}
	node0 := &TreeNode{
		Val: 0,
	}
	node4 := &TreeNode{
		Val: 4,
	}
	node7 := &TreeNode{
		Val: 7,
	}
	node9 := &TreeNode{
		Val: 9,
	}
	node3 := &TreeNode{
		Val: 3,
	}
	node5 := &TreeNode{
		Val: 5,
	}
	node6.Left = node2
	node6.Right = node8
	node2.Left = node0
	node2.Right = node4
	node8.Left = node7
	node8.Right = node9
	node4.Left = node3
	node4.Right = node5

	return node6
}

func BuildNode() *Node {
	root := &Node{
		Val: 1,
	}
	child1 := &Node{Val: 3, Children: []*Node{&Node{Val: 5}, {Val: 6}}}
	root.Children = append(root.Children, child1, &Node{Val: 2}, &Node{Val: 4})
	return root
}

func Entry() {
	//t := Build()
	//fmt.Println(largestValues(t))
	//CBuild()
	root := Build()
	dfsSearch(root, 4)
	for _, v := range result {
		fmt.Println(v.Val)
	}
	//fmt.Println(stack)
}

func BuildOrderTree(data []int) *TreeNode {
	var tree *TreeNode
	for _, d := range data {
		tree = insertNodeToTreeNoRec(tree, d)
	}

	return tree
}

func insetNodeToTree(t *TreeNode, val int) *TreeNode {
	if t == nil {
		return &TreeNode{Val: val}
	}
	if val < t.Val {
		t.Left = insetNodeToTree(t.Left, val)
	} else {
		t.Right = insetNodeToTree(t.Right, val)
	}
	return t
}

func insertNodeToTreeNoRec(t *TreeNode, val int) *TreeNode {
	newNode := &TreeNode{Val: val}
	if t == nil {
		return newNode
	}
	cur := t
	for cur != nil {
		if val < cur.Val {
			if cur.Left == nil {
				cur.Left = newNode
				break
			} else {
				cur = cur.Left
			}
		} else {
			if cur.Right == nil {
				cur.Right = newNode
				break
			} else {
				cur = cur.Right
			}
		}
	}
	return t
}
