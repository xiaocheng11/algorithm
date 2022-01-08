package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//构造最大二叉树:在数组中，最大值的左边是左子树，右边是右子树，本身就是头节点
func constructMaximumBinaryTree(nums []int) *TreeNode {
	return build(nums, 0, len(nums)-1)
}

func build(nums []int, l, r int) *TreeNode {
	if l > r {
		return nil
	}
	//找到最大值
	maxIndex := -1
	max := math.MinInt64
	for i := l; i <= r; i++ {
		if nums[i] > max {
			maxIndex = i
			max = nums[i]
		}
	}

	root := &TreeNode{
		Val: max,
	}

	//找到最大值后,找左边的跟右边的
	root.Left = build(nums, l, maxIndex-1)
	root.Right = build(nums, maxIndex+1, r)
	return root
}

func main() {

}
