package _5

//请实现一个函数，把字符串 s 中的每个空格替换成"%20"
func replaceSpace(s string) string {
	s1 := ""
	for i := 0; i < len(s); i++ {
		if string(s[i]) == " " {
			s2 := "%20"
			s1 += s2
			continue
		}
		s1 += string(s[i])
	}
	return s1
}
