package main

import "fmt"

func main() {
	fmt.Println(FindIndex("onesssstwo", "sssstwo"))
}
//给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从 0 开始)。如果不存在，则返回 -1。
//思路：将后面那个数组，每次比较，有一次不相同，就退出
func FindIndex(haystack, needle string) int {
	if len(haystack) == 0 || len(needle) == 0 {
		return -1
	}
	//先遍历大的字符串
	for i := 0; i < len(haystack) - len(needle) + 1; i++ {
		//这个是遍历nuddle这个字符串
		for j := 0; j < len(needle); j++ {
			if haystack[i + j] != needle[j] {
				//说明两者不相等，退出
				break
			}
			//说明两者相等
			if j == len(needle) - 1 {
				return i
			}
		}
	}
	return -1

}