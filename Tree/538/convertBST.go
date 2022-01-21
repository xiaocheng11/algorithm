package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//给出二叉 搜索 树的根节点，该树的节点值各不相同，请你将其转换为累加树（Greater Sum Tree），
//使每个节点 node的新值等于原树中大于或等于node.val的值之和。

func convertBST(root *TreeNode) *TreeNode {
   //就是从后面开始累加
	sum := 0
	var build func(root *TreeNode)
	build = func(root *TreeNode) {
		if root == nil {
			return
		}
		//将递归顺序倒过来
		build(root.Right)
		sum += root.Val
		root.Val = sum
		build(root.Left)
	}
	build(root)
	return root
}

func main() {

}
