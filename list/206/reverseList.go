package _06

type ListNode struct {
	Val  int
	Next *ListNode
}

//给你单链表的头节点 head ，请你反转链表，并返回反转后的链表。
//思路：给两个辅助结点，一个保存前面的，一个保存后面的结点
func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	//开始循环链表
	for head != nil {
		//保存后面的结点
		tmp := head.Next
		head.Next = pre
		// 保存当前head.Next节点，防止重新赋值后被覆盖
		// 一轮之后状态：nil<-1 2->03->4
		//              prev   head
		//这步是将原来的给接上去了
	   pre = head
	   //向下移动
	   head = tmp
	}
	return pre
}
