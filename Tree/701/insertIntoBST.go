package main

//给定二叉搜索树（BST）的根节点root和要插入树中的值value，将值插入二叉搜索树。
//返回插入后二叉搜索树的根节点。 输入数据 保证 ，新值和原始二叉搜索树中的任意节点值都不同。

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func insertIntoBST(root *TreeNode, val int) *TreeNode {
    if root == nil {
    	//说明找到了可以插入的位置
    	root = &TreeNode{Val: val}
    	return root
	}
	//然后去判断是大于还是小于
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	}else {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}
// DFS查找插入位置
func insertIntoBST1(root *TreeNode, val int) *TreeNode {
	if root == nil {
		root = &TreeNode{Val: val}
		return root
	}
	if root.Val > val {
		root.Left = insertIntoBST(root.Left, val)
	} else {
		root.Right = insertIntoBST(root.Right, val)
	}
	return root
}
func main() {

}
