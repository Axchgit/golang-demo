package main

import "fmt"

func main() {

	fmt.Println("start")

	var a *int

	b := 123

	a = &b

	fmt.Println(a, b)
}
