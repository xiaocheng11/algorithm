package main

import "fmt"

//插入排序：思想，将数组第一个看作有序，后面的看作无序的，每次从无序的拿一个，插入前面有序的去
func InsertionSort(array []int) []int {

	//第一个循环是要插入的数的下标
	for i := 1; i < len(array); i++ {
		preIndex := i - 1 //每次都是从无序数组的前一个
		current := array[i] //要交换的值
		//终止条件，当不大于的时候，或者到了最前面
		for preIndex >= 0 && array[preIndex] > current {
			array[preIndex + 1] = array[preIndex]
			preIndex--
		}
		//然后交换下标后一个的两个值
		array[preIndex + 1] = current
	}
	return array
}


func main() {
	arr := []int{1,5,8,9,7,6,10,-1,3,-2}
	arr = InsertionSort(arr)
	fmt.Println(arr)
}
