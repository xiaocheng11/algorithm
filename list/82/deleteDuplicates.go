package _2

type ListNode struct {
	Val  int
	Next *ListNode
}
//给定一个已排序的链表的头 head ， 删除原始链表中所有重复数字的节点，只留下不同的数字 。返回 已排序的链表 。
//思路：可能删除头节点，所以用哑巴节点做辅助节点

func deleteDuplicates(head *ListNode) *ListNode {
	//判断是否存在值
	if head == nil {
		return head
	}
	dummy := &ListNode{Val: 0}
	dummy.Next = head
	head = dummy
	//然后开始循环
	var rmVal int
	for head.Next != nil && head.Next.Next != nil {
		//这说明两个节点的值相等，需要删除
		if head.Next.Val == head.Next.Next.Val {
			//需要一个存放值，然后往下走
			rmVal = head.Next.Val
			for head.Next != nil && head.Next.Val == rmVal {
				head.Next = head.Next.Next
			}
		}else {
			head = head.Next
		}
	}
	return dummy.Next
}