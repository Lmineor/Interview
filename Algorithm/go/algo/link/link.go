package link

import "fmt"

func revertLinkInKPair(node *ListNode, k int) *ListNode {
	if Length(node) <= k {
		return revertLink(node)
	}
	var flag int = 1
	prev := new(ListNode)
	prev.Next = node
	prevCopy := prev
	start := node
	end := node
	for end != nil {
		if flag == k {
			nextSubLinkStart := end.Next // next subLink start
			end.Next = nil               // let it be a true subLink

			subLinkStart := start
			slice := revertLink(subLinkStart)
			prev.Next = slice

			// now start is end, prev is new subLink's prev
			prev = start
			end = nextSubLinkStart
			start = nextSubLinkStart
			flag = 1
		} else {
			flag++
			end = end.Next
		}
	}

	// link the remaining link
	prev.Next = start

	return prevCopy.Next
}

func Length(link *ListNode) int {
	var l int = 0
	for link != nil {
		l++
		fmt.Println(link.Val)
		link = link.Next
	}
	return l
}

func revertLink(node *ListNode) (newNode *ListNode) {
	if node == nil {
		return
	}
	next := node.Next
	current := node
	current.Next = nil
	for next != nil {
		tmp := next.Next
		next.Next = current
		current = next
		next = tmp
	}
	return current
}

func deleteItemInLink(node *ListNode, val int) *ListNode {
	if node == nil {
		return node
	}
	if node.Val == val {
		return node.Next
	}
	hair := new(ListNode)
	hair.Next = node
	for node != nil && node.Next != nil {
		if node.Next.Val == val {
			node.Next = node.Next.Next
		} else {
			node = node.Next
		}
	}
	return hair.Next
}

func merge2Links(link1, link2 *ListNode) *ListNode {
	l1 := Length(link1)
	l2 := Length(link2)
	var shorter, longer *ListNode
	if l1 < l2 {
		shorter = link1
		longer = link2
	} else {
		shorter = link2
		longer = link1
	}
	newOne := new(ListNode)
	pivNew := newOne
	piv1 := shorter
	piv2 := longer
	for piv1 != nil && piv2 != nil {
		if piv1.Val <= piv2.Val {
			pivNew.Next = piv1
			piv1 = piv1.Next
		} else {
			pivNew.Next = piv2
			piv2 = piv2.Next
		}
		pivNew = pivNew.Next
	}
	if piv1 != nil {
		pivNew.Next = piv1
	}
	if piv2 != nil {
		pivNew.Next = piv2
	}

	return newOne.Next
}

func findCross(link1, link2 *ListNode) *ListNode {
	fmt.Println("111")
	l1 := Length(link1)
	fmt.Println("111")
	l2 := Length(link2)
	var shorter, longer *ListNode
	var gap int
	if l1 < l2 {
		shorter = link1
		longer = link2
		gap = l2 - l1
	} else {
		shorter = link2
		longer = link1
		gap = l1 - l2
	}
	piv1 := longer
	for i := 0; i < gap; i++ {

		piv1 = piv1.Next
	}
	piv2 := shorter
	for piv1 != nil && piv2 != nil {
		fmt.Println("111")
		if piv1 == piv2 {
			return piv1
		} else {
			piv1 = piv1.Next
			piv2 = piv2.Next
		}
	}
	return &ListNode{}
}
