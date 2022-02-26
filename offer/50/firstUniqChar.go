package _0

//在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。
func firstUniqChar(s string) byte {
	if len(s) == 0 {
		return ' '
	}
    //利用hash
	ans := make([]int,26)
	for _, ch := range s {
		ans[ch - 'a']++
	}
	for i, ch := range s {
		if ans[ch - 'a'] == 1 {
			return s[i]
		}
	}
	return ' '
}
