package search_for_range

//给定一个包含 n 个整数的排序数组，找出给定目标值 target 的起始和结束位置。 如果目标值不在数组中，则返回[-1, -1]
//思路：核心点就是找第一个 target 的索引，和最后一个 target 的索引，所以用两次二分搜索分别找第一次和最后一次的位置

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	//然后做两次二分查找，分别找到起始位置，和结束位置
	result := make([]int, 2)
	start := 0
	end := len(nums) - 1
	for start + 1 < end {
		mid := start + (end - start) / 2
		if nums[mid] > target {
			end = mid
		}else if nums[mid] < target {
			start = mid
		}else {
			//先找左边的临界值
			end = mid
		}
	}
	if nums[start] == target {
		result[0] = start
	}else if nums[end] == target {
		result[0] = end
	}else {
		result[0] = -1
		result[1] = -1
		return result
	}
	start = 0
	end = len(nums) - 1
	//然后找中止位置
	for start + 1 < end {
		mid := start + (end - start) / 2
		if nums[mid] > target {
			end = mid
		}else if nums[mid] < target {
			start = mid
		}else {
			//先找左边的临界值
			start = mid
		}
	}
	if nums[start] == target {
		result[1] = start
	}else if nums[end] == target {
		result[1] = end
	}else {
		result[0] = -1
		result[1] = -1
		return result
	}
	return result
}
