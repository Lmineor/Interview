package link

//给你两个单链表的头节点 headA 和 headB ，请你找出并返回两个单链表相交的起始节点。如果两个链表没有交点，返回 null 。
//思路：让长的先跑比短的长的步数
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headB == nil {
		return nil
	}
	if headA == nil {
		return nil
	}
	la := lengthListNode(headA)
	lb := lengthListNode(headB)
	var longer, shorter *ListNode
	var step int = 0
	if la > lb {
		longer, shorter = headA, headB
		step = la - lb
	} else {
		longer, shorter = headB, headA
		step = lb - la
	}
	for i := 0; i < step; i++ {
		longer = longer.Next
	}
	for longer != nil && shorter != nil {
		if longer == shorter {
			return longer
		}
		longer = longer.Next
		shorter = shorter.Next
	}
	return nil
}

func lengthListNode(head *ListNode) int {
	var l = 0
	for head != nil {
		l++
		head = head.Next
	}
	return l
}
