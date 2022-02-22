package _43

//给定一个单链表 L 的头节点 head ，单链表 L 表示为：
//
//L0 → L1 → … → Ln - 1 → Ln
//请将其重新排列后变为：
//
//L0 → Ln → L1 → Ln - 1 → L2 → Ln - 2 → …
//不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

type ListNode struct {
	Val  int
	Next *ListNode
}

//思路，先找到中间结点，然后反转后面的，在合并两链表
func reorderList(head *ListNode) {
	if head == nil {
		return
	}
	//找到中心结点
	middle := findMiddle(head)
	//然后反转该尾部链表
	tail := reverseList(middle.Next)
	middle.Next = nil
	//最后合并两链表
	head = mergeTwoLists(head, tail)
}

func findMiddle(head *ListNode) *ListNode {
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	return slow
}
func reverseList(head *ListNode) *ListNode {
	//要一个前结点
	var pver *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = pver
		pver = head
		head = tmp
	}
	return pver
}

func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	//要组成一个新链表
	newNode := &ListNode{Val: 0}
	head := newNode
	toggle := true
	for list2 != nil && list1 != nil {
		if toggle {
			head.Next = list1
			list1 = list1.Next
		} else {
			head.Next = list2
			list2 = list2.Next
		}
		toggle = !toggle
		head = head.Next
	}
	if list1 != nil {
		head.Next = list1
	}
	if list2 != nil {
		head.Next = list2
	}
	return newNode.Next
}