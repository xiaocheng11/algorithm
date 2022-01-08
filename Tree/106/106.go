package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//根据一棵树的中序遍历与后序遍历构造二叉树。
func buildTree(inorder []int, postorder []int) *TreeNode {
	return build(inorder, 0, len(inorder)-1, postorder, 0, len(postorder)-1)
}

//建树的代码
func build(inorder []int, inStart, inEnd int, postorder []int, postStart, postEnd int) *TreeNode {
	//终止条件
	if inStart > inEnd {
		return nil
	}
	//根据后序遍历的特点，最后一个为根节点
	rootCur := postorder[postEnd]

	//然后根据中序遍历的情况，找到中间节点
	index := 0
	for i := inStart; i <= inEnd; i++ {
		if inorder[i] == rootCur {
			index = i
		}
	}
	//找到左子树的长度
	leftSize := index - inStart

	//然后新建一个节点
	root := &TreeNode{
		Val: rootCur,
	}
	//左子树
	root.Left = build(inorder, inStart, index -1, postorder, postStart, postStart + leftSize - 1)
	//右子树
	root.Right = build(inorder, index + 1, inEnd, postorder,postStart + leftSize, postEnd - 1)
	return root
}
func main() {

}
