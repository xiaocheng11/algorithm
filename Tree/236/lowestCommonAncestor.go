package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。
//
//百度百科中最近公共祖先的定义为：“对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，
//满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。”

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	//往左数跟右树要信息
	if root == nil {
		return root
	}
	//如果根节点等于其中一个，则直接返回
	if p == root || q == root {
		//说明，这个root为一个祖先
		return root
	}
	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)
	//如果两个都不为空，说明两个子树在两边
	if left != nil && right != nil {
		return root
	}

	if left != nil {
		return left
	}

	if right != nil {
		return right
	}
	//否则没有公共祖先
	return nil
}
func main() {

}
