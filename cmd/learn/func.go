package main

import "fmt"

var sum = func(a int, b int) int {
	return a + b
}

func main() {
	fmt.Println("start")

	fmt.Println(sum(2, 100))
}
