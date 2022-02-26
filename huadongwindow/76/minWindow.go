package _6

import "math"

//给你一个字符串 s 、一个字符串 t 。返回 s 中涵盖 t 所有字符的最小子串。如果 s 中不存在涵盖 t 所有字符的子串，则返回空字符串 "" 。

func minWindow(s string, t string) string {
	//首先，先创建两个字典，一个是窗口字符，一个是要查询的字符
	// 保存滑动窗口字符集
	win := make(map[byte]int)
	// 保存需要的字符集
	need := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}
	//然后窗口
	left, right := 0, 0
	//比较的次数
	match := 0
	//起始位置和终点
	start := 0
	end := 0
	//记录最小字符串
	min := math.MaxInt64
	var c byte
	//开始移动窗口
	for right < len(s) {
		c = s[right]
		right++
		//在需要的字符集里面，添加到窗口字符集里面
		if need[c] != 0 {
			win[c]++
			if win[c] == need[c] {
				//如果当前字符的数量匹配需要的字符的数量，则match值+1
				match++
			}
		}
		// 当所有字符数量都匹配之后，开始缩紧窗口
		for match == len(need) {
			//先判断当前窗口是否是最小的
			if right-left < min {
				start = left
				end = right
				min = right - left
			}
			//收缩窗口
			c = s[left]
			left++
			// 左指针指向不在需要字符集则直接跳过
			if need[c] != 0 {
				// 左指针指向字符数量和需要的字符相等时，右移之后match值就不匹配则减一
				// 因为win里面的字符数可能比较多，如有10个A，但需要的字符数量可能为3
				// 所以在压死骆驼的最后一根稻草时，match才减一，这时候才跳出循环
				if win[c] == need[c] {
					match--
				}
				win[c]--
			}
		}
	}

	if min == math.MaxInt64 {
		return ""
	}
	return s[start:end]
}
