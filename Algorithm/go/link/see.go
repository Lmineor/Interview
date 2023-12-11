package link

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func buildLink(data []int) *Node {
	hair := new(Node)
	head := hair
	for _, d := range data {
		node := Node{Val: d}
		head.Next = &node
		head = head.Next
	}
	return hair.Next
}

func printLink(node *Node) {
	for node != nil {
		if node.Next == nil {
			fmt.Printf("%d", node.Val)
		} else {
			fmt.Printf("%d->", node.Val)
		}

		node = node.Next
	}
	fmt.Println()
}

func Entry() {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	//data := []int{1, 2, 3, 4, 5, 6, 7, 8}
	node := buildLink(data)
	//node = revertLink(node)
	//node = deleteItemInLink(node, 6)
	l := revertLinkInKPair(node, 4)
	printLink(l)
}
