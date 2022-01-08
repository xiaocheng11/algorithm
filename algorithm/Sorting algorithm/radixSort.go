package main

import "fmt"

// 基数排序
func BaseSort(data []int) []int {
	if len(data) < 2 {
		return data
	}
	//找最大值
	max := data[0]
	dataLen := len(data)
	for i := 1; i < dataLen; i++ {
		if data[i] > max {
			max = data[i]
		}
	}
	// 计算最大值的位数
	maxDigit := 0
	for max > 0 {
		max = max/10
		maxDigit++
	}

	// 定义每一轮的除数，1,10,100...
	divisor := 1;
	// 定义了10个桶，为了防止每一位都一样所以将每个桶的长度设为最大,与原数组大小相同
	bucket := [10][20]int{{0}}
	// 统计每个桶中实际存放的元素个数
	count := [10]int{0}
	// 获取元素中对应位上的数字，即装入那个桶
	var digit int
	// 经过maxDigit+1次装通操作，排序完成
	for i := 1; i <= maxDigit; i++ {
		for j := 0; j < dataLen; j++ {
			tmp := data[j]
			digit = (tmp / divisor) % 10
			bucket[digit][count[digit]] = tmp
			count[digit]++
		}
		// 被排序数组的下标
		k := 0
		// 从0到9号桶按照顺序取出
		for b := 0; b < 10; b++ {
			if count[b] == 0 {
				continue
			}
			for c := 0; c < count[b]; c++ {
				data[k] = bucket[b][c]
				k++
			}
			count[b] = 0
		}
		divisor = divisor * 10
	}
	return data
}
func main() {
	arr := []int{1,5,8,9,7,6,10}
	arr = BaseSort(arr)
	fmt.Println(arr)
}
