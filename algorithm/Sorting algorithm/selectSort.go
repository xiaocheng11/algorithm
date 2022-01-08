package main

import "fmt"

//选择排序：思想每次都选择一个小的或大的，然后与第一位交换
func SelectSort(array []int) []int {
	//第一次循环，是每次需要交换的下标
	for i := 0; i < len(array) - 1; i++ {
		index := i
		//第二次循环是要找到最小值的下标
		for j := i + 1 ; j < len(array); j++ {
			if array[j] < array[index] {
				index = j
			}
		}
		//如果两个下标不相等，才做交换动作。
		if index != i {
		array[i], array[index] = array[index], array[i]
		}
	}
	return array
}

func main()  {
	arr := []int{1,5,8,9,7,6,10,-1,3,-2}
	arr = SelectSort(arr)
	fmt.Println(arr)
}