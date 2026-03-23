package main

import "fmt"

func main() {

	fmt.Println(123)

	mp := map[int]string{
		1: "a",
		2: "b",
	}

	a := map[int]string{
		1: "123",
	}

	a[1] = "change"

	fmt.Println(a)

	aa := make(map[string]int, 10)
	aa["a"] = 123
	fmt.Println(aa)

	// 错误，引用类型，没有分配内存  panic: assignment to entry in nil map
	// var ab map[string]int
	// ab["123"] = 123
	// fmt.Println(ab)

	for index, value := range mp {
		fmt.Println(index, value)
	}

	fmt.Println(mp)

}
