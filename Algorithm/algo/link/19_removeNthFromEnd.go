package link

//给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
//
//进阶：你能尝试使用一趟扫描实现吗？
//
//示例 1：
//
//19.删除链表的倒数第N个节点
//
//输入：head = [1,2,3,4,5], n = 2 输出：[1,2,3,5] 示例 2：
//
//输入：head = [1], n = 1 输出：[] 示例 3：
//
//输入：head = [1,2], n = 1 输出：[1]
//
//#

//双指针的经典应用，如果要删除倒数第n个节点，让fast移动n步，然后让fast和slow同时移动，直到fast指向链表末尾。删掉slow所指向的节点就可以了。

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p := &ListNode{}
	p.Next = head
	fast, slow := p.Next, p.Next
	pre, next := p, slow.Next
	for i := 1; i <= n && fast != nil; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		pre = slow
		slow = slow.Next
		next = slow.Next
	}
	if fast == nil{
		pre.Next = next
	}
	return p.Next
}
