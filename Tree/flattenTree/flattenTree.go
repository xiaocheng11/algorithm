package main

type TreeNode struct {
	value int
	left  *TreeNode
	right *TreeNode
}

//将二叉树变成链表：就是在后序遍历的地方做逻辑
func flattenTree(head *TreeNode) {
	if head == nil {
		return
	}
	flattenTree(head.left)
	flattenTree(head.right)
	//找到左树跟右树
	left := head.left
	right := head.right
    //将左树置为右树
	head.right = left
	head.left = nil

	//然后将右树放在左数的右数
	p := head
	for p.right != nil {
		p = p.right
	}
	p.right = right
}

func main() {

}
