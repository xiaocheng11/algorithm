package _2


type ListNode struct {
	Val  int
	Next *ListNode
}
//给你单链表的头指针 head 和两个整数left 和 right ，其中left <= right 。
//请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
//先遍历到 m 处，翻转，再拼接后续，注意指针处理
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	// 思路：先遍历到m处，翻转，再拼接后续，注意指针处理
	// 输入: 1->2->3->4->5->NULL, m = 2, n = 4
	if head == nil {
		return head
	}
  //有可能是投结点,所以使用哑巴结点
	dummy := &ListNode{Val: 0}
	dummy.Next = head
	head = dummy
	//然后开始循环，找到要反转的位置
	// 最开始：0->1->2->3->4->5->nil
	var pre *ListNode
	var i = 0
	for i < left {
		pre = head
		head = head.Next
		i++
	}
	// 遍历之后： 1(pre)->2(head)->3->4->5->NULL
	var next *ListNode
	// 用于中间节点连接
	var mid = head
	//这里开始循环
	for head != nil && i <= right {
       tmp := head.Next
       head.Next = next
       next = head
       head = tmp
       i++
	}
	//反转完之后，需要拼接起来
	pre.Next = next //这是前面的
	mid.Next = head //这是后面的
	return dummy.Next

}
