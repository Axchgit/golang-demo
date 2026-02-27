package main

import "fmt"

func main() {
	a, b := 10, 15
	a++
	b--
	fmt.Println(a, b)

	var c bool = true
	d := false

	// 验证 || 和 && 的优先级
	// 如果 || 优先级高于 &&，则 (c || d && d) 等价于 ((c || d) && d) -> false
	// 如果 && 优先级高于 ||，则 (c || d && d) 等价于 (c || (d && d)) -> true
	fmt.Println(c || d && d) // 实际结果为 true， 说明 && 优先级高于 ||
}
