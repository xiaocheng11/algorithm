package main

//给定一个二叉树，判断它是否是高度平衡的二叉树。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	var maxDept func(root *TreeNode) int
	maxDept = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := maxDept(root.Left)
		right := maxDept(root.Right)

		//判断两边是否都为平衡二叉树，而且两边高度差不超过1
		if left == -1 || right == -1 || left > right+1 || right > left+1 {
			return -1
		}
		if left > right {
			return left + 1
		}
		return right + 1
	}
	if maxDept(root) == -1 {
		return false
	}
	return true
}
func main() {

}
