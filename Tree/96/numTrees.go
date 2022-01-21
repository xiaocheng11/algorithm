package main

//给你一个整数 n ，求恰由 n 个节点组成且节点值从 1 到 n 互不相同的 二叉搜索树 有多少种？返回满足题意的二叉搜索树的种数。
func numTrees(n int) int {
	//思路，穷举发，每个节点的由左子树和右子树的排列组合来d
	mode := make([]int, n+1)
	mode[0] = 1 //
	mode[1] = 1 //只有一个数的时候只有棵平衡二叉树
	//求后面节点数，因为有一颗作为节点数，所以需要-1
	for i := 2; i <= n; i++ {
		//这个循环就是来求每个节点有多少个平衡二叉树的个数
		for j := 0; j <= i - 1; j++ {
			mode[i] += mode[j] * mode[i-j-1]
		}
	}
	return mode[n]

}

func main() {

}
