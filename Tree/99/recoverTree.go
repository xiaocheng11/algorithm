package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
//给你二叉搜索树的根节点 root ，该树中的两个节点的值被错误地交换。请在不改变其结构的情况下，恢复这棵树。
func recoverTree(root *TreeNode) {
   //根据中序遍历的特点，可以知道搜索树是有序的，找出两个打乱顺序的值
	//first是代表第一个位置，second是第二个位置
	//pre则是要去寻找的位置
	first := &TreeNode{}
	second := &TreeNode{}
	pre := &TreeNode{
		Val: math.MinInt64,
	}
	//定义一个函数，可以在内部调用
	var midSort func(root *TreeNode)
	midSort = func(root *TreeNode) {
		if root == nil {
			return
		}
		midSort(root.Left)
		//在这里进行coding
		if root.Val < pre.Val {
			//说明无序了，将first赋值
			if first.Val == 0 {
				first = pre
			}
			//第二次找到无序的值
			second = root
		}
		//每次都要向下走
		pre = root
		midSort(root.Right)
	}
	midSort(root)
	//然后交换两者的值
	first.Val, second.Val = second.Val, first.Val
}

func main() {

}
