package main

import "fmt"

func main() {
	var str string = "undefined"
	fmt.Println(str)

	var b, c int = 1, 5
	fmt.Println(b*c)

	var d *int
	fmt.Println(d)

	f, g := "good", "boy"
	fmt.Println(f+"  "+g)
}