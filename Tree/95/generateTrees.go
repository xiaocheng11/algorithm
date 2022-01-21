package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//不同的搜索二叉树2：给你一个整数 n ，请你生成并返回所有由 n 个节点组成且节点值从 1 到 n 互不相同的不同 二叉搜索树 。
//可以按 任意顺序 返回答案。
func generateTrees(n int) []*TreeNode {
	//从左树和右树要信息
	if n == 0 {
		return nil
	}
	//找到有多少个平衡二叉树
	return build(1, n)
}

//求有多少种平衡二叉树
func build(start, end int) []*TreeNode {
	//中止条件就是开始区间大于中止区间
	if start > end {
		return []*TreeNode{nil}
	}
	res := make([]*TreeNode, 0)
	//然后就是找到左边，跟右边的平衡二叉树
	for i := start; i <= end; i++ {
		//左边的切片，跟右边的切边合在一起
		leftSlice := build(start, i-1)
		rightSlice := build(i+1, end)
		//然后合并
		for _, left := range leftSlice {
			for _, right := range rightSlice {
				//然后创建根节点
				cur := &TreeNode{i, nil, nil}
				cur.Left = left
				cur.Right = right
				//将当前树加入到切片去
                res = append(res,cur)
			}
		}
	}
	return res
}

func main() {

}
