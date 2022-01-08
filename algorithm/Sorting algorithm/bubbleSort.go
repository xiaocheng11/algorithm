package main

import "fmt"

//冒泡排序：思想：每次从头或者尾巴开始，两两比较，将大的放到后面去。
func BubbleSort(array []int) []int {
	//第一次循环，是要比较的次数
	for i := 0; i <= len(array)-1; i++ {
		//第二次循环，是每次比较的数字,每次都可以减少比较的次数
       for j := 0; j < len(array) - i - 1; j++ {
		   if array[j] > array[j + 1] {
               array[j], array[j+1] = array[j+1],array[j]
		   }
	   }
	}
	return array
}

func main()  {
	arr := []int{1,5,8,9,7,6,10,-1,3,-2}
	arr = BubbleSort(arr)
	fmt.Println(arr)
}