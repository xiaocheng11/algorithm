package main

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
//124. 二叉树中的最大路径和
//路径 被定义为一条从树中任意节点出发，沿父节点-子节点连接，达到任意节点的序列。同一个节点在一条路径序列中 至多出现一次 。
//该路径 至少包含一个 节点，且不一定经过根节点。
//路径和 是路径中各节点值的总和。
//给你一个二叉树的根节点 root ，返回其 最大路径和 。
func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32
    //用来求最大值的函数
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := max(0,dfs(root.Left))
		right := max(0,dfs(root.Right))

		innerMaxSum := left + root.Val + right
		maxSum = max(maxSum, innerMaxSum)
		return root.Val + max(left, right) //只能选择一条路径，不能两头都选
	}
	dfs(root)
	return maxSum
}

func max(a, b int) int {
	if a > b { return a }
	return b
}

func main() {

}
