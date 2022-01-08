package main

import "strconv"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//给定一棵二叉树，返回所有重复的子树。对于同一类的重复子树，你只需要返回其中任意一棵的根结点即可。
//两棵树重复是指它们具有相同的结构以及相同的结点值。
func findDuplicateSubtrees(root *TreeNode) []*TreeNode {
	//用一个map记录该树出现过几次
	myMap := map[string]int{}
	//切片记录有哪些重复的树
    res := make([]*TreeNode, 0)
    build(root, myMap, &res)
    return res
}

//遍历，用深度遍历，将树序列化
func build(root *TreeNode, myMap map[string]int, res *[]*TreeNode) string {
	//如果节点为空，返回#
     if root == nil {
		 return "#"
	 }
	 //后序遍历
	 leftString := build(root.Left, myMap, res)
	 rightString := build(root.Right, myMap, res)

	 //然后将整棵树的信息拼接起来
	 rootString := leftString + rightString + strconv.Itoa(root.Val)

	 //然后查看map里面是否有这个值
	 if v,ok := myMap[rootString]; ok {
	 	if v == 1 {
			//取值，说明已经有相同的树了
			*res = append(*res, root)
		}
	 	myMap[rootString]++
	 }else {
		 myMap[rootString]++
	 }
	 return rootString
}


func main() {

}
