package main

import "fmt"

func main() {
	fmt.Println('A')

	var a [5]int
	a = [5]int{1, 2, 3}

	ab := [...]int{1, 2, 4}
	ab[2] = 2

	ac := new([4]int)
	ad := *ac

	ac[0] = 1 // 这里ac[0]是数组指针的语法糖，其实等级与 (*ac)[0] = 1
	(*ac)[1] = 6

	ae := 1
	var af *int
	af = &ae

	fmt.Println(a, ab, ac, ad, af)
}
