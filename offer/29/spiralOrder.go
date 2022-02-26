package _9

//输入一个矩阵，按照从外向里以顺时针的顺序依次打印出每一个数字。
func spiralOrder(matrix [][]int) []int {
	ans := make([]int, 0)
	if matrix == nil || len(matrix) == 0 || len(matrix[0]) == 0{
		return ans
	}
	//定义四个变量，分边代表上下左右边界
	r1, r2 := 0, len(matrix)-1
	c1, c2 := 0, len(matrix[0])-1
	for r1 <= r2 && c1 <= c2 {
		//先进行最上面一层的打印，从左至右
		for i := c1; i <= c2; i++ {
			ans = append(ans, matrix[r1][i])
		}
		//然后就是最右边的打印，从上至下
		for i := r1 + 1; i <= r2; i++ {
			ans = append(ans, matrix[i][c2])
		}
		//然后就是最下面的打印，从右至左，只有当r1 ！= r2 时才能
		if r1 != r2 {
			for i := c2 - 1; i >= c1; i-- {
				ans = append(ans, matrix[r2][i])
			}
		}
		//然后就是最左边的，从下至上,同理，两者不相等的时候才进行打印，否者会重复
		if c1 != c2 {
			for i := r2 - 1; i > r1; i-- {
				ans = append(ans, matrix[i][c1])
			}
		}
		//打印完后，要进行换行
		r1++
		r2--
		c1++
		c2--
	}
	return ans
}
