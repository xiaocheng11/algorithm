package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//给定一个二叉搜索树的根节点 root ，和一个整数 k ，请你设计一个算法查找其中第 k 个最小元素（从 1 开始计数）。
func kthSmallest(root *TreeNode, k int) int {
	res := 0
	len := 0
	var build func(root *TreeNode, k int)
	build = func(root *TreeNode, k int) {
		//判断条件
		if root == nil {
			return
		}
		//然后调用递归
		build(root.Left, k)
		//这是退出条件
		len++
		if len == k {
			res = root.Val
			return
		}
		build(root.Right, k)
	}
	build(root, k)
	return res
}

func main() {

}
