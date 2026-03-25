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
	fmt.Println(aa["111"])

	ab := make(map[int]string, 10)
	ab[1] = "123"
	fmt.Println(ab)

	if _, exit := ab[11]; exit {
		fmt.Println("11 exists")
	} else {
		fmt.Println(11)
	}

	delete(ab, 1)

	fmt.Println(ab)

	for i := 0; i < 10; i++ {
		ab[i] = "nihao:" + fmt.Sprintf("%d", i)
	}

	fmt.Println(ab[1])
	clear(ab)
	fmt.Println(ab[1])
	fmt.Println(ab)
}
