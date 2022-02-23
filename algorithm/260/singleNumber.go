package _60

//给定一个整数数组 nums，其中恰好有两个元素只出现一次，其余所有元素均出现两次。
//找出只出现一次的那两个元素。你可以按 任意顺序 返回答案。

//思路：利用位运算，之后找到最后一个1的位置，开始进行位运算
func singleNumber(nums []int) []int {
	sum := 0
	for _, num := range nums {
		sum = sum ^ num
	}
	result := []int{sum, sum}
	// 去掉末尾的1后异或diff就得到最后一个1的位置
	sum = (sum & (sum - 1)) ^ sum
	for i := 0; i < len(nums); i++ {
		if sum & nums[i] == 0 {
			result[0] = result[0]^nums[i]
		}else {
			result[1] ^= nums[i]
		}
	}
	return result
}
