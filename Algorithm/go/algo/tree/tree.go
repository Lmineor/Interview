package tree

import "fmt"

type Tree struct {
	Val   int
	Left  *Tree
	Right *Tree
}

func Build() *Tree {
	root := &Tree{
		Val: 1,
	}
	root.Left = &Tree{Val: 2, Left: &Tree{Val: 4}, Right: &Tree{Val: 5}}
	root.Right = &Tree{Val: 3, Left: &Tree{Val: 6}, Right: &Tree{Val: 7}}
	return root
}

func Entry() {
	t := Build()
	fmt.Println(postOrderTraversalNoRec(t))
}

func BuildOrderTree(data []int) *Tree {
	var tree *Tree
	for _, d := range data {
		tree = insertNodeToTreeNoRec(tree, d)
	}

	return tree
}

func insetNodeToTree(t *Tree, val int) *Tree {
	if t == nil {
		return &Tree{Val: val}
	}
	if val < t.Val {
		t.Left = insetNodeToTree(t.Left, val)
	} else {
		t.Right = insetNodeToTree(t.Right, val)
	}
	return t
}

func insertNodeToTreeNoRec(t *Tree, val int) *Tree {
	newNode := &Tree{Val: val}
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
