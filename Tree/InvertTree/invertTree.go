package main

import "fmt"

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

//翻转二叉树：考点，就是将子树的左子节点与右子节点交换
func invertTree(head *TreeNode) {
     if head == nil {
		 return
	 }
	 //在这里交换左子跟右子的节点
	 head.left, head.right = head.right, head.left
	 invertTree(head.left)
	 invertTree(head.right)
}

//递归序遍历二叉树
func rangeTree(head *TreeNode) {
	if head == nil {
		return
	}
	//这一片是前序遍历，根左右
	//fmt.Println(head.value)
	rangeTree(head.left)
	//这一片是中序遍历，左根右
	//fmt.Println(head.value)
	rangeTree(head.right)
	//这一片是后序遍历，左右根
	fmt.Println(head.value)
}
func main() {
    a := &TreeNode{
    	value: 10,
	}
	b := &TreeNode{
		value: 11,
	}
	c := &TreeNode{
		value: 12,
	}
	d := &TreeNode{
		value: 13,
	}
	a.left = b
	a.right = c
	b.left = d
	rangeTree(a)
	invertTree(a)
	rangeTree(a)
}
