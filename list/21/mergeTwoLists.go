package _1

type ListNode struct {
	Val  int
	Next *ListNode
}

//将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	newNode := &ListNode{Val: 0}
	head := newNode
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			head.Next = list1
			list1 = list1.Next
		} else {
			head.Next = list2
			list2 = list2.Next
		}
		//头节点也向下移动
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
