package main

import "fmt"

//归并排序：利用分治的思想，将数组分为左边跟右边，
//然后将两边有序的起来就是有序数组
//递归函数
func MergeSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}
	middle := len(array) / 2
	left := array[:middle]
	right := array[middle:]
	return merge(MergeSort(left), MergeSort(right))
}
//合并函数
func merge(left, right []int) []int {
	//开辟出一个新的数组，然后比较俩个数组，每次将最小的放进去
	var result []int
	//然后循环两个数组,两个都不能为空
	for len(left) != 0 && len(right) != 0 {
		if left[0] > right[0] {
			result = append(result, right[0])
			right = right[1:]
		}else {
			result = append(result, left[0])
			left = left[1:]
		}
	}

	//如果左边的数组不为空，将其放入result
	if len(left) != 0 {
		result = append(result, left...)
	}
	//相同，右边也一样
	if len(right) != 0 {
		result = append(result, right...)
	}
	return result
}
func main() {
	arr := []int{1,5,8,9,7,6,10,-1,3,-2}
	arr = MergeSort(arr)
	fmt.Println(arr)
}
