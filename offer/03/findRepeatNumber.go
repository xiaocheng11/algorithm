package _3

//找出数组中重复的数字。
//在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。
//数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。

func findRepeatNumber(nums []int) int {
	num := make(map[int]int)
	var n int
	for i := 0; i < len(nums); i++ {
		num[nums[i]]++
		if num[nums[i]] > 1 {
			n = nums[i]
			return n
		}
	}
	return n
}

//提升
func findRepeatNumber2(nums []int) int {
	//利用数组所在的位置，进行判断
	i := 0
	for i < len(nums) {
		//如果位置相同的话就不用换了
		if nums[i] == i {
			i++
			continue
		}

		if nums[nums[i]] == nums[i] {
			return nums[i]
		}
		nums[nums[i]], nums[i] = nums[i], nums[nums[i]]
	}
	return -1
}
