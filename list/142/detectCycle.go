package _42

type ListNode struct {
	Val  int
	Next *ListNode
}
//给定一个链表的头节点  head ，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。
//思路：快慢指针
func detectCycle(head *ListNode) *ListNode {
	// 思路：快慢指针，快慢相遇之后，慢指针回到头，快慢指针步调一致一起移动，相遇点即为入环点
	if head == nil {
		return head
	}
	fast := head.Next
	slow := head

	for fast != nil && fast.Next != nil {
		if fast == slow {
			// 慢指针重新从头开始移动，快指针从第一次相交点下一个节点开始移动
			fast = head
			slow = slow.Next // 注意
			// 比较指针对象（不要比对指针Val值）
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return slow
		}
		fast = fast.Next.Next
		slow = slow.Next
	}
	return nil
}
