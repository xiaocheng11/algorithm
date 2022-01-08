package main

import "fmt"

//希尔排序：思想，首先先确定步频，然后比较相差步频个数的数字
//然后将步频减少
func ShellSort(array []int) []int {
	//第一个循环，确定步频
	for grep := len(array) / 2; grep > 0; grep = grep / 2 {
         //第二循环，找要循环的个数，就是插入排序了
		for i := grep; i < len(array); i++ {
			preIndex := i - grep
			current := array[i]
			//判断是否要退出循环
			for preIndex >= 0 && array[preIndex] > current {
				array[preIndex + grep] = array[preIndex]
				preIndex = preIndex - grep
			}
			//交换两个值
			array[preIndex + grep] = current
		}
	}
	return array
}
func main() {
	arr := []int{1,5,8,9,7,6,10,-1,3,-2}
	arr = ShellSort(arr)
	fmt.Println(arr)
}
