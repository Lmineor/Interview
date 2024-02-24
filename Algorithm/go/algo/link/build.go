package link

func buildLink(data []int) *ListNode {
	hair := new(ListNode)
	head := hair
	for _, d := range data {
		node := ListNode{Val: d}
		head.Next = &node
		head = head.Next
	}
	return hair.Next
}

func buildCrossLink(data1, data2 []int) (node1 *ListNode, node2 *ListNode) {
	var shorter, longer *ListNode
	var gap int
	if len(data1) <= len(data2) {
		shorter = buildLink(data1)
		longer = buildLink(data2)
		gap = len(data2) - len(data1)
	} else {
		longer = buildLink(data1)
		shorter = buildLink(data2)
		gap = len(data1) - len(data2)
	}
	pivS := shorter
	pivL := longer
	for i := 0; i < gap; i++ {
		pivL = pivL.Next
	}
	preS := pivS
	preL := pivL
	for pivS != nil && pivL != nil {
		if pivL.Val == pivS.Val {
			preL.Next = pivL
			preS.Next = pivL
			break
		}
		preS = pivS
		preL = pivL
		pivL = pivL.Next
		pivS = pivS.Next
	}
	return shorter, longer
}
