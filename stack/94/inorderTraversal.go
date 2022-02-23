package _4

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//给定一个二叉树的根节点 root ，返回它的 中序 遍历。

func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)
	//利用栈去实现
	if root == nil {
		return result
	}

	stack := make([]*TreeNode, 0)
	for root != nil || len(stack) > 0 {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		//然后取出
		val := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		result = append(result, val.Val)
		root = val.Right
	}
	return result

}
