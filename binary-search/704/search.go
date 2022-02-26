package _04

//给定一个n个元素有序的（升序）整型数组nums 和一个目标值target ，
//写一个函数搜索nums中的 target，如果目标值存在返回下标，否则返回 -1。
func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	start := 0
	end := len(nums) - 1
	for start <= end {
		mid := start + (end - start) / 2
		if nums[mid] == target {
			return mid
		}else if nums[mid] > target {
			end = mid - 1
		}else if nums[mid] < target {
			start = mid + 1
		}
	}
	return -1
}
