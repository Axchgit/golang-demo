package main

import "fmt"

func main() {
	a, b := 0, 1
	// fmt.Scanln(&a, &b)
	// fmt.Println(a + b)

	msg := fmt.Sprintf("a + b = %d", a+b)
	fmt.Println(msg)

	str := "abcdef"

	newStr := fmt.Sprintf("% x", str)
	fmt.Println(newStr)
}
