package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
//给你二叉树的根节点 root ，返回其节点值的 层序遍历 。 （即逐层地，从左到右访问所有节点）。
//思路：用一个队列记录一层的元素，然后扫描这一层元素添加下一层元素到队列（一个数进去出来一次，所以复杂度 O(logN)）
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
