package main

import "fmt"

//快速排序：思想每次找一个基准值，将小于它的，放左边，等于的放中间，大于的放右边，然后两边继续递归
func QuickSort(arr []int, left, right int) []int {
	//退出条件
	if left < right {
        partitionIndex := partition(arr, left, right)
		QuickSort(arr, left, partitionIndex - 1) //左边的递归
		QuickSort(arr, partitionIndex + 1, right) //右边的递归
	}
	return arr
}
//用来将数组
func partition(arr []int, left, right int) int {
   //将第一个作为基准值
	pivot := left
	index := pivot + 1 //开始查找的下标
	for i := index; i <= right; i++ {
		//说明i这个位置上的小于基准，要交换
		if arr[pivot] >= arr[i] {
             arr[index], arr[i] = arr[i], arr[index]
             index ++
		}
	}
	//将小于值的
	arr[pivot], arr[index - 1] = arr[index - 1], arr[pivot]
	//返回这个基准值下标
	return index - 1
}
func main() {
	arr := []int{1,5,8,9,7,6,10,-1,3,-2}
	arr = quickSort(arr, 0, len(arr) - 1)
	fmt.Println(arr)
}

//快排
func quickSort(arr []int, left, right int) []int {
   if left < right {
   	//寻找基准值，然后通过基准值去快排
	   partitionIndex := partition2(arr, left, right)
	   quickSort(arr, left, partitionIndex - 1) //左边的递归
	   quickSort(arr, partitionIndex + 1, right) //右边的递归
   }
   return arr
}
//用来将数组进行分边
func partition2(arr []int, left, right int) int {
	//将第一个作为基准值
	pivot := left
	//这个是要去遍历的下标
	index := pivot + 1
	for i := index; i <= right; i++ {
         if arr[pivot] >= arr[i] {
         	//小于则说明要交换位置
         	arr[i], arr[index] = arr[index], arr[i]
         	index ++
		 }
	}
	//退出循环后，需要将基准值交换一下位置
	arr[pivot], arr[index - 1] = arr[index - 1], arr[pivot]
	//返回这个基准值下标
	return index - 1

}