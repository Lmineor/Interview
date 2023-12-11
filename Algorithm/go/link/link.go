package link

func revertLinkInKPair(node *Node, k int) *Node {
	if Length(node) <= k {
		return revertLink(node)
	}
	var flag int = 1
	prev := new(Node)
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

func Length(link *Node) int {
	var l int = 0
	for link != nil {
		l++
		link = link.Next
	}
	return l
}

func revertLink(node *Node) (newNode *Node) {
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

func deleteItemInLink(node *Node, val int) *Node {
	if node == nil {
		return node
	}
	if node.Val == val {
		return node.Next
	}
	hair := new(Node)
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
