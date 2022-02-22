package _48

type ListNode struct {
	Val  int
	Next *ListNode
}

//给你链表的头结点 head ，请将其按 升序 排列并返回 排序后的链表 。
func sortList(head *ListNode) *ListNode {
	//思路：使用归并排序
	return mergeSort(head)
}

//归并算法
func mergeSort(head *ListNode) *ListNode {
	//退出条件就是
	if head == nil || head.Next == nil {
		return head
	}
	//先找到中心点
	middle := findMiddle(head)
	tail := middle.Next
	middle.Next = nil
	//然后合并左边跟右边的信息
	left := mergeSort(head)
	right := mergeSort(tail)
	return mergeTwoLists(left, right)

}

//寻找中心位置
func findMiddle(head *ListNode) *ListNode {
	//用快慢指针
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

//合并两链表
func mergeTwoLists(list1, list2 *ListNode) *ListNode {
	//要组成一个新链表
	newNode := &ListNode{Val: 0}
	head := newNode
	for list2 != nil && list1 != nil {
		if list1.Val < list2.Val {
			head.Next = list1
			list1 = list1.Next
		} else {
			head.Next = list2
			list2 = list2.Next
		}
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
