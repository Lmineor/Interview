package link

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func printLink(node *ListNode) {
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
	data1 := []int{1,2}
	//data2 := []int{1, 2, 3, 4, 5}
	////data := []int{1, 2, 3, 4, 5, 6, 7, 8}
	node1 := buildLink(data1)
	////node2 := buildLink(data2)
	////node = revertLink(node)
	////node = deleteItemInLink(node, 6)
	////l := revertLinkInKPair(node, 4)
	////l := merge2Links(node1, node2)
	//l1, l2 := buildCrossLink(data1, data2)
	////printLink(l1)
	////printLink(l2)
	//node := findCross(l1, l2)
	//if node != nil {
	//	fmt.Println(node.Val)
	//}

	//ll := Constructor()
	//fmt.Println( ll.Get(1))
	//ll.AddAtHead(1)
	//ll.AddAtTail(3)
	//
	//fmt.Println(ll.ele)
	//ll.AddAtIndex(1,2)
	//fmt.Println(ll.ele)
	//ll.DeleteAtIndex(1)
	//fmt.Println(ll.ele)

	l := removeNthFromEnd(node1, 1)
	printLink(l)
}
