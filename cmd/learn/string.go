package main

import "fmt"

func main() {
	fmt.Println('a')
	fmt.Print(`
		343243243
		ewrsedrfdsf
		fdsfd
	`)
	a := "abcdefg "
	fmt.Println(a[1])
	fmt.Println(string(a[1]))
	fmt.Println(a[:2])
	fmt.Println(string(a[:2]))
	fmt.Println(string(a))
}
