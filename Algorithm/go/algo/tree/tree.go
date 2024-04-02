package tree

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Build() *TreeNode {
	root := &TreeNode{
		Val: 1,
	}
	root.Left = &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}
	root.Right = &TreeNode{Val: 3, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 7}}
	return root
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
	t := BuildNode()
	fmt.Println(levelOrderN(t))
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
