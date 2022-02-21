package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//给你二叉树的根节点 root ，返回其节点值 自底向上的层序遍历 。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）
//思路：在层级遍历的基础上，翻转一下结果即可
func levelOrderBottom(root *TreeNode) [][]int {
	result := levelOrder(root)
	// 翻转结果
	reverse(result)
	return result

}
//然后一个反转函数
func reverse(nums [][]int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
func levelOrder(root *TreeNode) [][]int {
	result := make([][]int,0)
	if root == nil {
		return result
	}
	//然后创建一个队列，将节点放进去
	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	//然后循环队列
	for len(queue) > 0 {
		//用来存放的list
		list := make([]int, 0)
		//这个是用来判断当前层数
		l := len(queue)
		for i := 0; i < l; i++ {
			//出队列
			level := queue[0]
			queue = queue[1:]
			//然后判断有无子树
			if level.Left != nil {
				queue = append(queue, level.Left)
			}
			if level.Right != nil {
				queue = append(queue, level.Right)
			}
			list = append(list, level.Val)
		}
		result = append(result, list)
	}
	return result
}
func main() {

}
