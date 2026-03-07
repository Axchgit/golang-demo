package main

import (
	"fmt"
	"slices"
)

func main() {
	fmt.Println('A')

	var a [5]int
	a = [5]int{1, 2, 3, 4, 5}

	ab := [...]int{1, 2, 4}
	ab[2] = 100

	ac := new([4]int)
	ad := *ac

	ac[0] = 1 // 这里ac[0]是数组指针的语法糖，其实等级与 (*ac)[0] = 1
	(*ac)[1] = 6

	dc := *ac

	ae := 1
	var af *int
	af = &ae
	da := a[:]
	da[0] = 99

	// a = append(a, 1) //append只支持 slice 参数
	fmt.Println(a, da, a, &a[0], &a[1], &a[2], &da[3], &a[4])
	// 对切片使用append后，切片与原数组不再是 同片个内存 存储
	da = append(da, 0)
	da[0] = 101
	fmt.Println(a, da, &a[0], &da[0])

	db := slices.Clone(a[:])

	fmt.Println(&a, ab, ac, ad, af, da, cap(da), cap(db), dc)
}
