package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//根据前序和中序，返回二叉树的根节点
func buildTree(preorder []int, inorder []int) *TreeNode {
	return build(preorder, 0, len(preorder)-1, inorder, 0, len(inorder)-1)
}
//建树的过程
func build(preorder []int, preStart, preEnd int, inorder []int, inStart, inEnd int) *TreeNode {
	if preStart > preEnd {
		return nil;
	}
	//根据前序的特点，头节点就是第一个
	rootVal := preorder[preStart]
    index := 0
    //找到在中序遍历里的头节点
	for i := inStart; i <= inEnd; i++ {
		if inorder[i] == rootVal {
			index = i
		}
	}

	 leftSize := index - inStart;
	//找到根节点，补上左节点，右节点
	root := &TreeNode{
		Val: rootVal,
	}
	root.Left = build(preorder, preStart + 1, preStart + leftSize, inorder, inStart, index - 1)
	root.Right = build(preorder,preStart + leftSize + 1, preEnd, inorder, index + 1, inEnd)
	return root
}

func main() {

}
