package _3

type ListNode struct {
	Val  int
	Next *ListNode
}

//给定一个已排序的链表的头 head ， 删除所有重复的元素，使每个元素只出现一次 。返回 已排序的链表 。
//思路就是一个一个删下去
func deleteDuplicates(head *ListNode) *ListNode {
   current := head
   for current != nil {
   	//然后在这判断是否会相等
   	for current.Next != nil && current.Val == current.Next.Val {
   		//这就跳过当前节点
   		current.Next = current.Next.Next
	}
	//直到没有找到后才往下走
	current = current.Next
   }
   return head
}
