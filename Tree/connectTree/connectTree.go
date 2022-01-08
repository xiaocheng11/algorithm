package main

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
	next  *TreeNode
}

//填充二叉树节点的右侧指针:就是将每个左子树指向右子树，右子树指向空
func connect(head *TreeNode) {
	if head == nil {
		return
	}
	//辅助函数，用来链接左子树和右子树
	connectTree(head.left, head.right)
}

//用来链接左数和右数
func connectTree(left, right *TreeNode) {
	if left == nil || right == nil {
		return
	}
	//链接
	left.next = right
	//左树跟右树都要链接
	connectTree(left.left, left.right)
	connectTree(right.left, right.right)
	// 连接跨越父节点的两个子节点
	connectTree(left.right, right.left)
}
func main() {

}
