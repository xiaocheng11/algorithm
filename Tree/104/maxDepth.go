package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//给定一个二叉树，找出其最大深度。
//二叉树的深度为根节点到最远叶子节点的最长路径上的节点数
//就是从子树要信息，左边跟右边的
func maxDepth(root *TreeNode) int {
    if root == nil {
    	return 0
	}
	//然后比较左树，跟右树
	left := maxDepth(root.Left)
	right := maxDepth(root.Right)
	if left > right {
		return left + 1
	}
	return right + 1
}
func main() {

}
