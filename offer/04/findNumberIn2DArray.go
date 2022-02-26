package _4

//在一个 n * m 的二维数组中，每一行都按照从左到右递增的顺序排序，每一列都按照从上到下递增的顺序排序。
//请完成一个高效的函数，输入这样的一个二维数组和一个整数，判断数组中是否含有该整数。

func findNumberIn2DArray(matrix [][]int, target int) bool {
	if matrix == nil || len(matrix) == 0|| len(matrix[0]) == 0 {
		return false
	}
	//从右上角开始走，左边的比它小，下边的比他大
	x := 0
	l := len(matrix[0]) - 1
	for x <= len(matrix) - 1&& l >= 0 {
		if matrix[x][l] == target {
			return true
		}else if matrix[x][l] > target {
			l--
		}else {
			x++
		}
	}
	return false
}
