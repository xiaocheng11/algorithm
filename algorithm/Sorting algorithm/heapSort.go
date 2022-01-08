package main

import "fmt"

//堆排序：有大顶堆跟小顶堆，首先先将数组给弄成大顶堆，
//然后每次交换顶堆的值与末尾交换

func HeadSort(array []int) []int {
	lenth := len(array)
	//第一步，先将数据写成大顶堆
    heapInsert(array, lenth)
    //然后就是取出数据，做heapify
    for i := lenth - 1; i >= 0; i-- {
    	//交换值
    	array[i],array[0] = array[0],array[i]
    	//然后长度-1
    	lenth--
    	//
    	heapify(array, 0, lenth)
	}
	return array
}
//形成大顶堆
func heapInsert(arr []int, arrLen int) {
	for i := arrLen / 2; i >= 0; i-- {
		heapify(arr, i, arrLen)
	}
}

//新建一个值的时候，从新建立大顶堆过程
func heapify(array []int, i, length int) {
	left := 2*i + 1  //左子节点
	right := 2*i + 2 //右子节点
	largest := i     //最大值的下标
	//左子节点大的时候
	if left < length && array[left] > array[largest] {
		largest = left
	}
	//右子节点大的时候
	if right < length && array[right] > array[largest] {
		largest = right
	}
	//交换i和largest的值
	if largest != i {
		array[i], array[largest] = array[largest], array[i]
		//然后再做子节点的heapify
		heapify(array, largest, length)
	}

}

func main() {
	arr := []int{1,5,8,9,7,6,10,-1,3,-2}
	arr = HeadSort(arr)
	fmt.Println(arr)
}
