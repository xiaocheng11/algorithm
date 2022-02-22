package _6

//给你一个链表的头节点 head 和一个特定值 x ，请你对链表进行分隔，使得所有 小于 x 的节点都出现在 大于或等于 x 的节点之前。
//
//你应当 保留 两个分区中每个节点的初始相对位置。

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
//思路，找到大于的重新放入一个链表中，最后在合并起来
	if head == nil {
		return head
	}
	headDummy := &ListNode{Val: 0}
	tailDummy := &ListNode{Val: 0}
	tail := tailDummy
	headDummy.Next = head
	head = headDummy
	for head.Next != nil {
		if head.Next.Val < x {
			head = head.Next
		}else {
			//这里需要将链表移出去
			tail.Next = head.Next
			head.Next = head.Next.Next
			tail = tail.Next
		}
	}
	//需要将tail的尾巴除干净
	tail.Next = nil
	head.Next = tailDummy.Next
	return headDummy.Next
}
