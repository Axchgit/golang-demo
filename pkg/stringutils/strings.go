package stringutils

import "fmt"

// Reverse 反转字符串
func Reverse(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		fmt.Println(string('交' + 1), string(r[i]), string(r[j]))
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// Upper 首字母大写
func Upper(s string) string {
	if s == "" {
		return s
	}
	r := []rune(s)
	if r[0] >= 'a' && r[0] <= 'z' {
		r[0] -= 32
	}
	return string(r)
}

// Lower 首字母小写
func Lower(s string) string {
	if s == "" {
		return s
	}
	r := []rune(s)
	if r[0] >= 'A' && r[0] <= 'Z' {
		r[0] += 32
	}
	return string(r)
}
